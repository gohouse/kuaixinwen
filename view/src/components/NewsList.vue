<template>
    <div>
        <scroller style="margin-top: 3rem;padding: 0 1rem;"
                  :on-refresh="refresh"
                  :on-infinite="infinite"
                  ref="my_scroller">
            <Timeline>
                <TimelineItem color="green">
                    <Icon type="trophy" slot="dot"></Icon>
                    <p><span>测试版本上线, 有问题尽管说, 虽然我不一定解决<^_^></span></p>
                    <p><span>手机访问,体验更佳,项目地址直通车: <a href="https://github.com/gohouse/kuaixinwen">github</a></span></p>
                </TimelineItem>
                <TimelineItem v-for="item in items">{{item.content}}</TimelineItem>
            </Timeline>
        </scroller>
    </div>
</template>

<script>
    export default {
        data () {
            return {
                page: 1,
                limit: 10,
                items: []
            }
        },

        mounted() {
            this.getNewsList()
        },

        methods: {
            refresh(done) {
                setTimeout(() => {
                    this.getNewsList(done)
                }, 1500)
            },

            infinite(done) {

                setTimeout(() => {
                    this.getNewsList(done)
                }, 1500)
            },

            getNewsList(done){
                this.axios.get("/getnewslist", {params:{"page":this.page,"limit":this.limit}}).then(response => {
                    if (response.data.status != 200) {
                        this.util.checkResponse(response.data);
                        return;
                    }
                    if (response.data.data == "" || response.data.data == null) {
                        done(true)
                        this.$Message.info("暂无更多新闻 !");
                        return
                    }
                    this.items = this.items.concat(response.data.data)
                    this.page = this.page+1
                    done()
                }, response => {
                    // error callback
                    this.$Message.error("请求错误, 请重试2 !");
                })
            }
        }
    }
</script>

<style>
    .my-scroller .pull-to-refresh-layer .spinner-holder
    {
        visibility: hidden;
    }

    .my-scroller .pull-to-refresh-layer {
        background-image: url(http://qianka.b0.upaiyun.com/images/4f013b6bc7d96fc347f416ad3673f937.png);
        background-repeat: no-repeat;
        background-position: center;
        background-size: 40px 40px;

        opacity: 0;
        -webkit-transform: scale(1);
        transform: scale(1);

        transition: all .15s linear;
        -webkit-transition: all .15s linear;
    }

    .my-scroller .pull-to-refresh-layer.active {
        -webkit-transform: scale(1.35);
        transform: scale(1.35);
        opacity: 1;
    }
</style>