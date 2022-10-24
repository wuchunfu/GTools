<template>
    <div style="padding: 0 150px;">
        <n-space vertical>
            <n-input type="text" placeholder="今日事今日毕" @keyup.enter="addItem()" v-model:value="todoItem"
                class="item_input" round />
        </n-space>
        <div>
            <n-gradient-text :size="40" type="success" class="total_todo">
                待办事项: {{total}}
            </n-gradient-text>
        </div>
        <div class="todo_detail">
            <p class="todo_detail_undo">未完成: {{undo}}</p>
            <p class="todo_detail_done">已完成: {{done}}</p>
            <p style="display: flex; align-items: center;">
                <n-button text style="font-size: 30px; margin-left: 15px;">
                    <n-icon>
                        <list-icon />
                    </n-icon>
                </n-button>
                <n-button text style="font-size: 30px; margin-left: 15px;">
                    <n-icon>
                        <hide-done-icon />
                    </n-icon>
                </n-button>
            </p>

            <n-progress style="width: 40px; line-height: 30px; margin-left: auto; margin-right: 20px;" class="progress"
                type="circle" :show-indicator="false" status="success" :percentage="todopPrcent" :stroke-width="12" />
        </div>
        <div>
            <n-list hoverable clickable>
                <n-list-item v-for="item, index in todoList">
                    <div @contextmenu.prevent="rightClick(item)" @dblclick.native="editItem(item)"
                        style="height: 40px; line-height: 40px; display: flex; align-items: center; justify-content: left;">
                        <n-checkbox size="large" :checked="item.done"
                            @update:checked="handleSelectionChange(item, index)" />
                        <p :style="item.done ? doneItemTitleStyle : todoItemTitleStyle">{{item.title}}</p>
                        <n-button text style="font-size: 30px; margin-left: 50px;" v-if="item.hasContent">
                            <n-icon :depth="item.done ? 4 : 2" color="#ff6347">
                                <content-icon />
                            </n-icon>
                        </n-button>
                        <n-time :time="new Date(item.date)" format="yyyy-MM-dd"
                            :style="item.done  ? doneItemDateStyle : todoItemDateStyle" />
                        <!-- <p :style="item.done  ? doneItemDateStyle : todoItemDateStyle">
                            {{item.date}}</p> -->
                    </div>
                </n-list-item>
            </n-list>
            <n-divider />
        </div>
        <n-back-top :right="100" />
    </div>
</template>
<script>
import {
    AlbumsOutline as ListIcon,
    DocumentText as ContentIcon,
    EyeOutline as HideDoneIcon
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
        HideDoneIcon
    },
    data() {
        return {
            app: window.go.gtools.App,
            todoItem: null,
            total: 0,
            undo: 0,
            done: 0,
            todopPrcent: 3,
            todoList: [],
            todoItemDateStyle: 'margin-left: auto; margin-right: 20px;font-weight: 500; font-size: large; color: #87ceeb;',
            doneItemDateStyle: 'margin-left: auto; margin-right: 20px;font-weight: 500; font-size: large; color: #D3D3D3;',
            doneItemTitleStyle: 'margin-left: 20px; font-size: large; font-weight: 600; text-decoration: line-through; color: #D3D3D3',
            todoItemTitleStyle: 'margin-left: 20px; font-size: large; font-weight: 600;',
        }
    },
    mounted() {
        this.app.GetTodoList().then(res => {
            if (res.code == 200) {
                this.todoList = res.data
            } else {
                message.error(res.msg)
            }
        })
    },
    methods: {
        addItem() {
            this.app.AddTodoItem({ title: this.todoItem }).then(res => {
                console.log(res.data);
                if (res.code == 200) {
                    this.todoList = res.data
                } else {
                    message.error(res.msg)
                }

            })
        },
        handleSelectionChange(item, index) {
            item.done = !item.done
            console.log(item);
            console.log(index);
        },
        tableRowClassName(val) {
            if (val.row.done === true) {
                return 'success-row'
            } else {
                return 'warning-row'
            }
        },
        rightClick(item) {
            let _this = this
            dialog.warning({
                title: "警告",
                content: "确定删除待办事项: " + item.title,
                positiveText: "确定",
                negativeText: "取消",
                onPositiveClick: () => {
                    _this.app.DelTodoItem({id: item.id}).then(res => {
                        if(res.code == 200) {
                            message.success("已删除");
                            _this.todoList = res.data
                        } else {
                            message.error(res.msg)
                        }
                    })
                    
                },
                onNegativeClick: () => {
                }
            });
            console.log(item);
        },
        editItem(item) {
            console.log(item);
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