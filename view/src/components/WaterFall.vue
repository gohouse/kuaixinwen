<template>
    <div class="books" v-scroll = "loadMore">
        <!--<Timeline>-->
            <!--<TimelineItem color="green">-->
                <!--<Icon type="trophy" slot="dot"></Icon>-->
                <!--<p><span>测试版本上线, 有问题尽管说, 虽然我不一定解决<^_^></span></p>-->
                <!--<p><span>项目地址直通车: <a href="https://github.com/gohouse/kuaixinwen">github</a></span></p>-->
            <!--</TimelineItem>-->
            <!--<TimelineItem v-for="item in data" :key="item.id">{{item.content}}</TimelineItem>-->
        <!--</Timeline>-->
        <scroller :on-infinite="infinite" ref="my_scroller">
            <div style="height: 44px;"></div>
            <div v-for="(item, index) in items"
                 class="row1" :class="{'grey-bg': index % 2 == 0}">
                {{ item }}
            </div>
        </scroller>
    </div>
</template>
<style>
    .row1 {
        width: 100%;
        height: 50px;
        padding: 10px 0;
        font-size: 16px;
        line-height: 30px;
        text-align: center;
        color: #444;
        background-color: #fff;
    }
    .grey-bg {
        background-color: #eee;
    }
</style>
<script>
    export default {
        data () {
            return {
                page: 1,
                limit: 5,
                data: [
                    {
                        "id": 36,
                        "status": 1,
                        "updated_at": null,
                        "content": "新闻啊新闻新闻啊新闻新闻啊新闻新闻啊新闻新闻啊新闻新闻啊新闻新闻啊新闻新闻啊新闻新闻啊新闻新闻啊新闻新闻啊新闻新闻啊新闻新闻啊新闻新闻啊新闻新闻啊新闻新闻啊新闻新闻啊新闻新闻啊新闻新闻啊新闻新闻啊新闻新闻啊新闻新闻啊新闻新闻啊新闻新闻啊新闻新闻啊新闻新闻啊新闻新闻啊新闻新闻啊新闻",
                        "created_at": null,
                    }
                ],
                items: [],
                sw: true,
            }
        },
        mounted () {
            for (let i = 1; i <= 20; i++) {
                this.items.push(i + ' - keep walking, be 2 with you.')
            }
            this.top = 1
            this.bottom = 20
        },
        created(){
        },
        methods: {
            infinite(done) {
//                if (this.bottom >= 30) {
//                    setTimeout(() => {
//                        done()
//                    }, 1500)
//                    return;
//                }
                setTimeout(() => {
                    let start = this.bottom + 1
                    for (let i = start; i < start + 10; i++) {
                        this.items.push(i + ' - keep walking, be 2 with you.')
                    }
                    this.bottom = this.bottom + 10;
                    setTimeout(() => {
                        done()
                    })
                }, 1500)
            },
            NewsList(){
                this.axios.get("/news/getlist", {params:{"page":this.page,"limit":this.limit}}).then(response => {
                    if (response.data.status != 200) {
                        this.util.checkResponse(response.data);
                        return;
                    }
                    this.data = response.data.data
                }, response => {
                    // error callback
                    this.$Message.error("请求错误, 请重试2 !");
                })
            }
        },
//        directives: {
//            scroll: {
//                bind: function (el, binding){
//                    window.addEventListener('scroll', ()=> {
//                        if(document.body.scrollTop + window.innerHeight >= el.clientHeight) {
//                            binding.value.call(this.NewsList())
//                        }
//                    })
//                }
//            }
//        }
    }
</script>
