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
                        <cash-icon />
                    </n-icon>
                </n-button>
                <n-button text style="font-size: 30px; margin-left: 15px;">
                    <n-icon>
                        <list-icon />
                    </n-icon>
                </n-button>
            </p>

            <n-progress style="width: 40px; line-height: 30px; margin-left: auto; margin-right: 20px;" class="progress"
                type="circle" :show-indicator="false" status="success" :percentage="todopPrcent" :stroke-width="12" />
        </div>
        <div>
            <!-- <el-table :data="todoList" class="table">
                <el-table-column prop="done"  width="80">
                    <template #default="scope">
                        <el-checkbox v-model="scope.row.done" @change="handleSelectionChange(scope.row)"></el-checkbox>
                    </template>
                </el-table-column>
                <el-table-column prop="address"  show-overflow-tooltip/>
                <el-table-column prop="content" width="120"  show-overflow-tooltip/>
                <el-table-column prop="date"  width="120" />
            </el-table> -->
            <n-table :bordered="false" :single-line="false">
                <tbody>
                    <tr v-for="item, index in todoList">
                        <td>
                            <el-checkbox v-model="item.done" @change="handleSelectionChange(item)"></el-checkbox>
                        </td>
                        <td>{{item.date}}</td>
                        <td style="overflow: hidden;">{{item.address}}</td>
                    </tr>
                </tbody>
            </n-table>
            <n-list hoverable clickable>
                <n-list-item v-for="item, index in todoList">
                    <n-thing title="相见恨晚" content-style="margin-top: 10px;">
                        <template #description>
                            <n-space size="small" style="margin-top: 4px">
                                <n-tag :bordered="false" type="info" size="small">
                                    暑夜
                                </n-tag>
                                <n-tag :bordered="false" type="info" size="small">
                                    晚春
                                </n-tag>
                            </n-space>
                        </template>
                        奋勇呀然后休息呀<br>
                        完成你伟大的人生
                    </n-thing>
                </n-list-item>
            </n-list>
        </div>
        <n-back-top :right="100" />
    </div>
</template>
<script>
import {
    Location as CashIcon,
    ListCircle as ListIcon
} from "@vicons/ionicons5";
export default {
    components: {
        CashIcon,
        ListIcon
    },
    data() {
        return {
            todoItem: null,
            total: 0,
            undo: 0,
            done: 0,
            todopPrcent: 3,
            todoList: [
                {
                    done: false,
                    date: '2016-05-03',
                    name: 'Tom',
                    address: 'No. 189, Grove St, Los Angeles No. 189, Grove St, Los AngelesNo. 189, Grove St, Los AngelesNo. 189, Grove St, Los AngelesNo. 189, Grove St, Los Angeles',
                },
            ]
        }
    },
    mounted() {
        for (let index = 0; index < 1; index++) {
            this.todoList.push({
                done: true,
                date: '2016-05-01',
                name: 'Tom',
                address: 'No. 189, Grove St, Los Angeles',
            })
        }
    },
    methods: {
        addItem() {
            console.log(this.todoItem);
        },
        handleSelectionChange(val) {
            console.log(val);
            console.log(val.date);
        },
        tableRowClassName(val) {
            if (val.row.done === true) {
                return 'success-row'
            } else {
                return 'warning-row'
            }
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
    color: #C03F53;
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