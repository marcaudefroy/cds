import {Component, Input, ViewChild} from '@angular/core';
import {Router} from '@angular/router';
import {TranslateService} from '@ngx-translate/core';
import {cloneDeep} from 'lodash';
import {ModalTemplate, SuiModalService, TemplateModalConfig} from 'ng2-semantic-ui';
import {ActiveModal} from 'ng2-semantic-ui/dist';
import {Subscription} from 'rxjs';
import {PermissionValue} from '../../../../../model/permission.model';
import {Project} from '../../../../../model/project.model';
import {
    Workflow,
    WorkflowNode,
    WorkflowNodeHook,
    WorkflowNodeJoin,
    WorkflowNodeTrigger,
    WorkflowPipelineNameImpact
} from '../../../../../model/workflow.model';
import {PipelineStore} from '../../../../../service/pipeline/pipeline.store';
import {WorkflowCoreService} from '../../../../../service/workflow/workflow.core.service';
import {WorkflowEventStore} from '../../../../../service/workflow/workflow.event.store';
import {WorkflowStore} from '../../../../../service/workflow/workflow.store';
import {AutoUnsubscribe} from '../../../../../shared/decorator/autoUnsubscribe';
import {ToastService} from '../../../../../shared/toast/ToastService';
import {WorkflowNodeConditionsComponent} from '../../../../../shared/workflow/node/conditions/node.conditions.component';
import {WorkflowNodeContextComponent} from '../../../../../shared/workflow/node/context/workflow.node.context.component';
import {WorkflowDeleteNodeComponent} from '../../../../../shared/workflow/node/delete/workflow.node.delete.component';
import {WorkflowNodeHookFormComponent} from '../../../../../shared/workflow/node/hook/form/hook.form.component';
import {HookEvent} from '../../../../../shared/workflow/node/hook/hook.event';
import {WorkflowTriggerComponent} from '../../../../../shared/workflow/trigger/workflow.trigger.component';

@Component({
    selector: 'app-workflow-sidebar-edit-node',
    templateUrl: './workflow.sidebar.edit.node.component.html',
    styleUrls: ['./workflow.sidebar.edit.node.component.scss']
})
@AutoUnsubscribe()
export class WorkflowSidebarEditNodeComponent {

    // Project that contains the workflow
    @Input() project: Project;
    @Input() workflow: Workflow;

    nodeSub: Subscription;

    // Child component
    @ViewChild('workflowTrigger')
    workflowTrigger: WorkflowTriggerComponent;
    @ViewChild('workflowDeleteNode')
    workflowDeleteNode: WorkflowDeleteNodeComponent;
    @ViewChild('workflowContext')
    workflowContext: WorkflowNodeContextComponent;
    @ViewChild('workflowConditions')
    workflowConditions: WorkflowNodeConditionsComponent;
    @ViewChild('worklflowAddHook')
    worklflowAddHook: WorkflowNodeHookFormComponent;
    @ViewChild('nodeNameWarningModal')
    nodeNameWarningModal: ModalTemplate<boolean, boolean, void>;

    // Modal
    @ViewChild('nodeParentModal')
    nodeParentModal: ModalTemplate<boolean, boolean, void>;
    newParentNode: WorkflowNode;
    modalParentNode: ActiveModal<boolean, boolean, void>;
    newTrigger: WorkflowNodeTrigger = new WorkflowNodeTrigger();
    node: WorkflowNode;
    previousNodeName: string;
    pipelineSubscription: Subscription;
    displayInputName = false;
    loading = false;
    nameWarning: WorkflowPipelineNameImpact;
    permissionEnum = PermissionValue;

    constructor(private _workflowStore: WorkflowStore, private _translate: TranslateService, private _toast: ToastService,
                private _pipelineStore: PipelineStore, private _modalService: SuiModalService,
                private _router: Router,
                private _workflowCoreService: WorkflowCoreService, private _workflowEventStore: WorkflowEventStore) {
        this.nodeSub = this._workflowEventStore.selectedNode().subscribe(n => {
            if (n) {
                if (!this.displayInputName) {
                    this.previousNodeName = n.name
                }
            }
            this.node = n;
        });
    }

    canEdit(): boolean {
        return this.workflow.permission === PermissionValue.READ_WRITE_EXECUTE;
    }

    openAddHookModal(): void {
        if (this.canEdit() && this.worklflowAddHook) {
            this.worklflowAddHook.show();
        }
    }

    openTriggerModal(): void {
        if (!this.canEdit()) {
            return;
        }
        this.newTrigger = new WorkflowNodeTrigger();
        this.newTrigger.workflow_node_id = this.node.id;
        if (this.workflowTrigger) {
          this.workflowTrigger.show();
        }
    }

    openAddParentModal(): void {
        if (!this.canEdit()) {
            return;
        }
        this.newParentNode = new WorkflowNode();
        let tmpl = new TemplateModalConfig<boolean, boolean, void>(this.nodeParentModal);
        this.modalParentNode = this._modalService.open(tmpl);
    }

    addNewParentNode(): void {
        let workflowToUpdate = cloneDeep(this.workflow);
        let oldRoot = cloneDeep(this.workflow.root);
        workflowToUpdate.root = this.newParentNode;
        if (oldRoot.hooks) {
            this.newParentNode.hooks = oldRoot.hooks;
        }
        delete oldRoot.hooks;
        workflowToUpdate.root.triggers = new Array<WorkflowNodeTrigger>();
        let t = new WorkflowNodeTrigger();
        t.workflow_dest_node = oldRoot;
        workflowToUpdate.root.triggers.push(t);

        this.updateWorkflow(workflowToUpdate, this.modalParentNode);
    }

    openWarningModal(): void {
        let tmpl = new TemplateModalConfig<boolean, boolean, void>(this.nodeNameWarningModal);
        this._modalService.open(tmpl);
    }

    openDeleteNodeModal(): void {
        if (!this.canEdit()) {
            return;
        }
        if (this.workflowDeleteNode) {
            this.workflowDeleteNode.show();
        }
    }

    openEditContextModal(): void {
        let sub = this.pipelineSubscription =
            this._pipelineStore.getPipelines(this.project.key, this.node.pipeline_name).subscribe(pips => {
                if (pips.get(this.project.key + '-' + this.node.pipeline_name)) {
                    setTimeout(() => {
                        this.workflowContext.show();
                        sub.unsubscribe();
                    }, 100);
                }
            });
    }

    openEditRunConditions(): void {
        this.workflowConditions.show();
    }

    saveTrigger(): void {
        if (!this.canEdit()) {
            return;
        }
        let clonedWorkflow: Workflow = cloneDeep(this.workflow);
        let currentNode: WorkflowNode;
        if (clonedWorkflow.root.id === this.node.id) {
            currentNode = clonedWorkflow.root;
        } else {
            currentNode = Workflow.getNodeByID(this.node.id, clonedWorkflow);
        }

        if (!currentNode) {
            return;
        }

        if (!currentNode.triggers) {
            currentNode.triggers = new Array<WorkflowNodeTrigger>();
        }
        currentNode.triggers.push(cloneDeep(this.newTrigger));
        this.updateWorkflow(clonedWorkflow, this.workflowTrigger.modal);
    }

    updateNodeConditions(n: WorkflowNode): void {
        if (!this.canEdit()) {
            return;
        }
        let clonedWorkflow: Workflow = cloneDeep(this.workflow);
        let node = Workflow.getNodeByID(n.id, clonedWorkflow);
        if (!node) {
            return;
        }
        node.context.conditions = cloneDeep(n.context.conditions);
        this.updateWorkflow(clonedWorkflow, this.workflowConditions.modal);
    }

    updateNode(n: WorkflowNode): void {
        if (!this.canEdit()) {
            return;
        }
        let clonedWorkflow: Workflow = cloneDeep(this.workflow);
        let node = Workflow.getNodeByID(n.id, clonedWorkflow);
        if (!node) {
            return;
        }
        node.context = cloneDeep(n.context);
        delete node.context.application;
        delete node.context.environment;
        this.updateWorkflow(clonedWorkflow, this.workflowContext.modal);
    }

    deleteNode(b: string): void {
        if (!this.canEdit()) {
            return;
        }
        let clonedWorkflow: Workflow = cloneDeep(this.workflow);
        if (b === 'all') {
            // Delete Node with Child
            if (clonedWorkflow.root.id === this.node.id) {
                this.deleteWorkflow(clonedWorkflow, this.workflowDeleteNode.modal);
                return;
            } else {
                clonedWorkflow = Workflow.removeNodesInNotifications(clonedWorkflow, clonedWorkflow.root, this.node.id, false);

                clonedWorkflow.root.triggers.forEach((t, i) => {
                    this.removeNode(clonedWorkflow, this.node.id, t.workflow_dest_node, clonedWorkflow.root, i);
                });
                if (clonedWorkflow.joins) {
                    clonedWorkflow.joins.forEach(j => {
                        j.source_node_ref = j.source_node_ref.filter(id => {
                           return id !== this.node.id.toString();
                        });
                        if (j.triggers) {
                            j.triggers.forEach((t, i) => {
                                this.removeNodeFromJoin(clonedWorkflow, this.node.id, t.workflow_dest_node, j, i);
                            });
                        }
                    });
                }
            }
        } else if (b === 'only') {
            let ok = Workflow.removeNodeWithoutChild(clonedWorkflow, this.node);
            if (!ok) {
                this._toast.error('', this._translate.instant('workflow_node_remove_multiple_parent'));
                return;
            }
            clonedWorkflow = Workflow.removeNodeInNotifications(clonedWorkflow, this.node);
        }
        this.updateWorkflow(clonedWorkflow, this.workflowDeleteNode.modal);
    }

    removeNodeFromJoin(workflow: Workflow, id: number, node: WorkflowNode, parent: WorkflowNodeJoin, index: number) {
        if (!this.canEdit()) {
            return;
        }
        if (node.id === id) {
            parent.triggers.splice(index, 1);
        }
        if (node.triggers) {
            node.triggers.forEach((t, i) => {
                this.removeNode(workflow, id, t.workflow_dest_node, node, i);
            });
        }
    }

    removeNode(workflow: Workflow, id: number, node: WorkflowNode, parent: WorkflowNode, index: number): Workflow {
        if (!this.canEdit()) {
            return;
        }
        if (node.id === id) {
            parent.triggers.splice(index, 1);
            workflow = Workflow.removeNodeInNotifications(workflow, node);
        }
        if (node.triggers) {
            node.triggers.forEach((t, i) => {
                workflow = this.removeNode(workflow, id, t.workflow_dest_node, node, i);
            });
        }
        return workflow;
    }

    deleteWorkflow(w: Workflow, modal: ActiveModal<boolean, boolean, void>): void {
        this._workflowStore.deleteWorkflow(this.project.key, w).subscribe(() => {
            this._toast.success('', this._translate.instant('workflow_deleted'));
            modal.approve(true);
            this._workflowEventStore.unselectAll();
            this._router.navigate(['/project', this.project.key], { queryParams: { tab: 'workflows'}});
        });
    }

    updateWorkflow(w: Workflow, modal?: ActiveModal<boolean, boolean, void>): void {
        this.loading = true;
        this._workflowStore.updateWorkflow(this.project.key, w).subscribe(() => {
            this.loading = false;
            this._toast.success('', this._translate.instant('workflow_updated'));
            this._workflowEventStore.unselectAll();
            if (modal) {
                modal.approve(null);
            }
        }, () => {
            if (Array.isArray(this.node.hooks) && this.node.hooks.length) {
              this.node.hooks.pop();
            }
            this.loading = false;
        });
    }

    createJoin(): void {
        if (!this.canEdit()) {
            return;
        }
        if (!this.node.ref) {
            this.node.ref = this.node.id.toString();
        }
        let clonedWorkflow: Workflow = cloneDeep(this.workflow);
        if (!clonedWorkflow.joins) {
            clonedWorkflow.joins = new Array<WorkflowNodeJoin>();
        }
        let j = new WorkflowNodeJoin();
        j.source_node_ref.push(this.node.ref);
        clonedWorkflow.joins.push(j);
        this.updateWorkflow(clonedWorkflow);
    }

    openRenameArea(): void {
        if (!this.canEdit()) {
            return;
        }
        this.nameWarning = Workflow.getNodeNameImpact(this.workflow, this.node.name);
        this.displayInputName = true;
    }

    linkJoin(): void {
        if (!this.canEdit()) {
            return;
        }
        this._workflowCoreService.linkJoinEvent(this.node);
    }

    addHook(he: HookEvent): void {
        if (!this.canEdit()) {
            return;
        }
        if (!this.node.hooks) {
            this.node.hooks = new Array<WorkflowNodeHook>();
        }
        this.node.hooks.push(he.hook);
        this.updateWorkflow(this.workflow, this.worklflowAddHook.modal);
    }

    rename(): void {
        if (!this.canEdit()) {
            return;
        }
        this.updateNode(this.node);
    }
}
