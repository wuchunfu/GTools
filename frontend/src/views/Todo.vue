<template>
    <div style="padding: 0 120px;">
        <n-space vertical>
            <n-input type="text" placeholder="今日事今日毕" @keyup.enter="addItem()" v-model:value="todoItem"
                class="item_input" round />
        </n-space>
        <div>
            <n-gradient-text :size="40" type="success" class="total_todo">
                待办事项: {{ data.todonum }}
            </n-gradient-text>
        </div>
        <div class="todo_detail">
            <p class="todo_detail_undo">未完成: {{ data.todonum }}</p>
            <p class="todo_detail_done">已完成: {{ data.donenum }}</p>
            <p style="display: flex; align-items: center;">
                <n-button text style="font-size: 30px; margin-left: 15px;">
                    <n-icon>
                        <list-icon />
                    </n-icon>
                </n-button>
                <n-button text style="font-size: 30px; margin-left: 15px;" @click="showDoneList = !showDoneList">
                    <n-icon>
                        <show-done-icon />
                    </n-icon>
                </n-button>
            </p>

            <n-progress style="width: 40px; line-height: 30px; margin-left: auto; margin-right: 20px;" class="progress"
                type="circle" :show-indicator="false" status="success" :percentage="data.rate" :stroke-width="12" />
        </div>
        <div>
            <n-list hoverable clickable>
                <n-list-item v-for="item, index in data.todoList">
                    <div @contextmenu.prevent="delTodoItem(item)" @dblclick.native="showEditModal(item)"
                        style="height: 40px; line-height: 40px; display: flex; align-items: center; justify-content: left;">
                        <n-checkbox size="large" :checked="item.done ? true : false"
                            @update:checked="handleSelectionChange(item, index)" />
                        <n-ellipsis :style="item.done ? doneItemTitleStyle : todoItemTitleStyle">{{ item.title }}</n-ellipsis>
                        <n-button text style="font-size: 30px; margin-left: 50px;" v-if="item.hasContent">
                            <n-icon :depth="item.done ? 4 : 2" color="#ff6347">
                                <content-icon />
                            </n-icon>
                        </n-button>
                        <n-space style="margin-left: auto;">
                            <n-tag :type="levelStyle[item.level]">
                                {{levels[item.level].label}}
                            </n-tag>
                        </n-space>
                        <n-time :time="new Date(item.date)" format="yyyy-MM-dd"
                            :style="item.done ? doneItemDateStyle : todoItemDateStyle" />
                    </div>
                </n-list-item>
            </n-list>
            <n-divider title-placement="center" v-show="showDoneList">已完成</n-divider>
            <n-list hoverable clickable v-show="showDoneList">
                <n-list-item v-for="item, index in data.doneList">
                    <div @contextmenu.prevent="rightClick(item)" @dblclick.native="showEditModal(item)"
                        style="height: 40px; line-height: 40px; display: flex; align-items: center; justify-content: left;">
                        <n-checkbox size="large" :checked="item.done ? true : false"
                            @update:checked="handleSelectionChange(item, index)" />
                        <p :style="item.done ? doneItemTitleStyle : todoItemTitleStyle">{{ item.title }}</p>
                        <n-button text style="font-size: 30px; margin-left: 50px;" v-if="item.hasContent">
                            <n-icon :depth="item.done ? 4 : 2" color="#ff6347">
                                <content-icon />
                            </n-icon>
                        </n-button>
                        <n-time :time="new Date(item.date)" format="yyyy-MM-dd"
                            :style="item.done ? doneItemDateStyle : todoItemDateStyle" />
                    </div>
                </n-list-item>
            </n-list>
        </div>
        <n-modal v-model:show="showItemModal" transform-origin="center" :auto-focus="false" :mask-closable="false">
            <n-card style="width: 600px" title="编辑事项" :bordered="false" size="huge" role="dialog" aria-modal="true">
                <n-form ref="formRef" :model="editItem" :rules="rules" label-placement="left" label-width="auto"
                    require-mark-placement="right-hanging" size="medium" :style="{
                        maxWidth: '640px'
                    }">
                    <n-form-item label="事项" path="todoValue">
                        <n-input v-model:value="editItem.title" placeholder="请输入事项" />
                    </n-form-item>
                    <n-form-item label="标签" path="tagsValue" :show-label="true">
                        <n-dynamic-tags v-model:value="tags" />
                    </n-form-item>
                    <!-- <n-form-item label="日期" path="dateValue" :show-label="true">
                        <n-date-picker v-model:value="editItem.mdate" type="date" />
                    </n-form-item> -->
                    <n-form-item label="等级" path="importValue" :show-label="true">
                        <n-radio-group v-model:value="editItem.level" name="radiobuttongroup">
                            <n-radio-button v-for="level in levels" :key="level.value" :value="level.value"
                                :label="level.label" />
                        </n-radio-group>
                    </n-form-item>
                    <n-form-item label="详情" path="contentValue" :show-label="true">
                        <n-input v-model:value="editItem.content" type="textarea" placeholder="详情" />
                    </n-form-item>
                    <n-form-item style="display: flex; justify-content: center;">
                        <n-button attr-type="button" @click="delById(editItem.id)" type="error" secondary round>
                            删除
                        </n-button>
                        <n-button attr-type="button" @click="closeModal()" style="margin-left: 30px" round secondary>
                            取消
                        </n-button>
                        <n-button style="margin-left: 30px" attr-type="button" type="success" @click="updateTodoItem(editItem.title)" round secondary>
                            保存
                        </n-button>
                    </n-form-item>
                </n-form>
            </n-card>
        </n-modal>
        <n-back-top :right="100" />
    </div>
</template>
<script>
import {
    AlbumsOutline as ListIcon,
    DocumentText as ContentIcon,
    EyeOutline as ShowDoneIcon
} from "@vicons/ionicons5";
import { createDiscreteApi } from 'naive-ui'
// 脱离上下文的 API 引入消息提示框
const { message, dialog } = createDiscreteApi(
    ["message", "dialog"]
);

export default {
    components: {
        ListIcon,
        ContentIcon,
        ShowDoneIcon
    },
    data() {
        return {
            app: window.go.gtools.App,
            todoItem: null,
            tags: [],
            levels: [
                { value: 0, label: '不急' },
                { value: 1, label: '普通' },
                { value: 2, label: '重要' },
                { value: 3, label: '紧急' },
            ],
            levelStyle: [
                '',
                'info',
                'warning',
                'error'
            ],
            data: {
                todonum: 0,
                donenum: 0,
                rate: 0,
                todoList: [],
                doneList: []
            },
            showDoneList: false,
            todoItemDateStyle: 'margin-left: 20px; margin-right: 20px;font-weight: 500; font-size: large; color: #87ceeb;',
            doneItemDateStyle: 'margin-left: auto; margin-right: 20px;font-weight: 500; font-size: large; color: #D3D3D3;',
            doneItemTitleStyle: 'margin-left: 20px; font-size: large; font-weight: 600; text-decoration: line-through; color: #D3D3D3; max-width: 320px',
            todoItemTitleStyle: 'margin-left: 20px; font-size: large; font-weight: 600; max-width: 320px',
            showItemModal: false,
            editItem: null,
            rules: {
                todoValue: {
                    required: true,
                },
            }
        }
    },
    mounted() {
        this.getTodoList()
    },
    methods: {
        getTodoList() {
            this.app.GetTodoList().then(res => {
                if (res.code == 200) {
                    this.data = res.data
                    this.tags = []
                } else {
                    message.error(res.msg)
                }
            })
        },
        addItem() {
            if (this.todoItem == null || this.todoItem.trim() == '') {
                return
            }
            this.app.AddTodoItem({ title: this.todoItem, date: new Date() ,level: 1}).then(res => {
                if (res.code == 200) {
                    this.data = res.data
                    this.todoItem = null
                } else {
                    message.error(res.msg)
                }
            })
        },
        handleSelectionChange(item) {
            item.done = !item.done
            this.app.UpdateTodoItem(item).then(res => {
                if (res.code != 200) { message.error(res.msg) }
                else {
                    this.data = res.data
                }
            })
        },
        updateItem(item) {
            this.app.UpdateTodoItem(item).then(res => {
                if (res.code != 200) { message.error(res.msg) }
                else {
                    this.data = res.data
                }
            })
        },
        tableRowClassName(val) {
            if (val.row.done === true) {
                return 'success-row'
            } else {
                return 'warning-row'
            }
        },
        delTodoItem(item) {
            let _this = this
            dialog.warning({
                title: "警告",
                content: "确定删除待办事项: " + item.title,
                positiveText: "确定",
                negativeText: "取消",
                onPositiveClick: () => {
                    _this.app.DelTodoItem({ id: item.id }).then(res => {
                        if (res.code == 200) {
                            message.success("已删除");
                            _this.data = res.data
                        } else {
                            message.error(res.msg)
                        }
                    })
                },
                onNegativeClick: () => {

                }
            });
        },
        showEditModal(item) {
            item.mdate = parseInt(new Date(item.date).getTime())
            this.editItem = item
            if(this.editItem.tags != '') {
                this.tags = this.editItem.tags.split(',')
                console.log(this.tags);
            }
            this.showItemModal = true
        },
        closeModal() {
            this.getTodoList()
            this.showItemModal = false
        },
        updateTodoItem(title) {
            if (title == "") {
                message.error("内容不完整")
                return
            }
            if (this.editItem.content != '' && this.editItem.content != null) {
                this.editItem.hasContent = true
            }
            this.editItem.tags = this.tags.join(',')
            this.updateItem(this.editItem)
            this.showItemModal = false
            this.tags = []
        },
        delById(id) {
            this.showItemModal = false
            this.app.DelTodoItemById({id: id}).then(res => {
                if(res.code == 200) {
                    this.data = res.data
                    this.tags = []
                } else{
                    message.error(res.msg)
                }
            })
        }
    }
}
</script>
<style>
.item_input_div {
    width: 500px;
}

.item_input {
    padding-left: 10px;
    margin-top: 20px;
    font-size: 20px;
    height: 50px;
    line-height: 50px;
}

.total_todo {
    margin: 10px 0 10px 20px;
}

.todo_detail {
    border: solid 2px #E0E0E5;
    border-left: none;
    border-right: none;
    display: flex;
    align-items: center;
    justify-content: left;
}

.todo_detail_done {
    border: solid 2px #E0E0E5;
    border-top: none;
    height: 30px;
    line-height: 30px;
    border-bottom: none;
    padding: 0 15px;
    margin-left: 15px;
    font-size: 20px;
    font-weight: 500;
    color: #4B9D5F;
}

.todo_detail_undo {
    color: #DB7093;
    font-size: 20px;
    font-weight: 500;
}

.el-table .warning-row {
    --el-table-tr-bg-color: var(--el-color-warning-light-9);
}

.el-table .success-row {
    --el-table-tr-bg-color: var(--el-color-success-light-9);
}

.table {
    width: 100%;
    border-radius: 10px;
    margin-top: 10px;
    transition: all 5s, 2s linear 0;
}

/* .table:hover {
    box-shadow: 0px 0px 3px 3px #DEDFE0;
} */
</style>