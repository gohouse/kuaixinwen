<template>
    <div class="books timeline" v-scroll = "loadMore">
        <Timeline>
            <TimelineItem color="green">
                <Icon type="trophy" slot="dot"></Icon>
                <p><span>测试版本上线, 有问题尽管说, 虽然我不一定解决<^_^></span></p>
                <p><span>项目地址直通车: <a href="https://github.com/gohouse/kuaixinwen">github</a></span></p>
            </TimelineItem>
            <TimelineItem v-for="item in data">{{item.content}}</TimelineItem>
        </Timeline>
    </div>
</template>
<style>
    .book {
        display: flex;
        margin-bottom: 20px;
        padding: 20px;
        border-radius: 5px;
        overflow: hidden;
        background-color: #fff;
    }
    .book__img {
        margin-right: 20px;
    }
</style>
<script>
    export default {
        data () {
            return {
                books: [
                    {
                        "id":1,
                        "title":"title",
                        "summary":"asfdsfsdf",
//                        "image":"/static/images/1.jpg"
                    }
                ],
                page: 1,
                scrollDisable: false
            }
        },
        mounted () {
            const _this = this
            // 数据请求
        },
        methods: {
            loadMore () {
                const _this = this
                if (!this.scrollDisable) {
                    // 一开始加载，就将 scrollDisable 设置为 true，即使触发了多次 loadMore，都只会执行一次下面的代码
                    this.scrollDisable = true
                    // 数据请求
                }
            }
        },
        directives: {
            scroll: {
                bind: function (el, binding){
                    window.addEventListener('scroll', ()=> {
                        if(document.body.scrollTop + window.innerHeight >= el.clientHeight) {
                            binding.value.call(this)
                        }
                    })
                }
            }
        }
    }
</script>
