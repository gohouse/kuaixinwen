<template>
    <div>
        <br>
        <Row style="text-align: right;">
            <Button icon="plus-circled" @click="BtnNewsAdd" type="primary">新增</Button>
        </Row>
        <br>
        <!--<Table border :columns="columns" :data="data"></Table>-->
        <Alert v-for="(item,index) in data">
            <Row>
                {{item.content}}
            </Row>
            <Row style="text-align: right;">
                <Button type="primary" @click="show(index)" size="small">修改</Button>
                <Button type="error" @click="remove(index)" size="small">删除</Button>
            </Row>
        </Alert>


        <Row style="text-align: center;">
            <Page :current="currentPage" :page-size="pageSize" :total="total" @on-change="OnPageChange"></Page>

        </Row>
        <br>
        <br>

    </div>
</template>
<script>
    export default {
        data () {
            return {
                columns: [
                    {
                        title: 'id',
                        key: 'id',
                        width: 80,
                        render: (h, params) => {
                            return h('div', [
//                                h('Icon', {
//                                    props: {
//                                        type: 'person'
//                                    }
//                                }),
                                h('strong', params.row.id)
                            ]);
                        }
                    },
                    {
                        title: '内容',
                        key: 'content'
                    },
                    {
                        title: '创建时间',
                        key: 'created_at',
                        width: 170
                    },
                    {
                        title: 'Action',
                        key: 'action',
                        width: 150,
                        align: 'center',
                        fixed: 'right',
                        render: (h, params) => {
                            return h('div', [
                                h('Button', {
                                    props: {
                                        type: 'primary',
                                        size: 'small'
                                    },
                                    style: {
                                        marginRight: '5px'
                                    },
                                    on: {
                                        click: () => {
                                            this.show(params.index)
                                        }
                                    }
                                }, '修改'),
                                h('Button', {
                                    props: {
                                        type: 'error',
                                        size: 'small'
                                    },
                                    on: {
                                        click: () => {
                                            this.remove(params.index)
                                        }
                                    }
                                }, '删除')
                            ]);
                        }
                    }
                ],
                data: [
                    {
                        id: 'John Brown',
                        content: 'New York No. 1 Lake Park',
                        created_at: 18,
                    },
                    {
                        name: 'Jim Green',
                        age: 24,
                        address: 'London No. 1 Lake Park'
                    },
                    {
                        name: 'Joe Black',
                        age: 30,
                        address: 'Sydney No. 1 Lake Park'
                    },
                    {
                        name: 'Jon Snow',
                        age: 26,
                        address: 'Ottawa No. 2 Lake Park'
                    }
                ],
                current_id: null,
                current_content: null,
                current_index: null,
                currentPage: 1,
                total: 100,
                pageSize: 5
            }
        },
        created() {
            this.GetNewsList()
        },
        methods: {
            BtnNewsAdd () {
                this.$Modal.confirm({
                    title: '新增',
//                    content: '<p>Content of dialog</p><p>Content of dialog</p>',
                    okText: '新增',
                    cancelText: '取消',
                    onOk: () => {
                        // 执行修改请求
                        this.NewsAddOrEdit('add')
                    },
                    render: (h) => {
                        return h('div',[
                            h('br'),
                            h('Input', {
                                props: {
//                                    value: this.current_content,
                                    autofocus: true,
                                    placeholder: '输入新闻信息...',
                                    type: "textarea",
                                    rows: "8"
                                },
                                on: {
                                    input: (val) => {
                                        this.current_content = val;
                                    }
                                }
                            })
                        ])
                    }
                });
            },
            show (index) {
                console.log(index)
                this.current_id = this.data[index].id
                this.current_content = this.data[index].content
                this.current_index = index

                this.$Modal.confirm({
                    title: '修改',
//                    content: '<p>Content of dialog</p><p>Content of dialog</p>',
                    okText: '修改',
                    cancelText: '取消',
                    onOk: () => {
                        // 执行修改请求
                        this.NewsAddOrEdit('edit')
                    },
                    onCancel: () => {
                        this._resetCurrentData()
                    },
                    render: (h) => {
                        return h('div',[
                            h('br'),
                            h('Input', {
                                props: {
                                    value: this.current_content,
                                    autofocus: true,
                                    placeholder: 'Please enter your name...',
                                    type: "textarea",
                                    rows: "8"
                                },
                                on: {
                                    input: (val) => {
                                        this.current_content = val;
                                    }
                                }
                            })
                        ])
                    }
                });
            },
            remove(index){
                this.current_id = this.data[index].id
                this.$Modal.confirm({
                    title: '删除',
                    content: '是否删除???',
                    okText: '删除',
                    cancelText: '取消',
                    onOk: () => {
                        // 执行修改请求
                        this.NewsDel()
                    },
                    onCancel: () => {
                        this._resetCurrentData()
                    }
                });
            },
            GetNewsList(){
                this.axios.get("/admin/getnewslist", {params:{"page":this.currentPage,"limit":this.pageSize}}).then(response => {
                    if (response.data.status != 200) {
                        this.util.checkResponse(response.data.msg);
                        return;
                    }
                    this.data = response.data.data
                    this.total = response.data.ext.total
                }, response => {
                    // error callback
                    this.$Message.error("请求错误, 请重试2 !");
                })
            },
            NewsAddOrEdit(act){
                // 执行请求
                this.axios.post("/admin/newsaddoredit", {"id":this.current_id,"content":this.current_content}).then(response => {
                    if (response.data.status != 200) {
                        this.util.checkResponse(response.data.msg);
                    } else {
                        this.data[this.current_index].content = this.current_content
                    }
                    this._resetCurrentData()
                }, response => {
                    // error callback
                    this.$Message.error("请求错误, 请重试3 !");
                })
                if (act=='add') this.data = this.GetNewsList()
            },
            _resetCurrentData(){
                this.current_id = null
                this.current_content = null
                this.current_index = null
            },
            NewsDel(){
                // 执行请求
                this.axios.post("/admin/newsdel", {"id":this.current_id}).then(response => {
                    if (response.data.status != 200) {
                        this.util.checkResponse(response.data.msg);
                    } else {
                        this.data = this.GetNewsList()
                    }
                    this._resetCurrentData()
                }, response => {
                    // error callback
                    this.$Message.error("请求错误, 请重试4 !");
                })
            },
            OnPageChange(index){
                this.currentPage = index
                this.data = this.GetNewsList()
            }
        }
    }
</script>
