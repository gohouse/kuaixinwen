<template>
    <div class="books" v-scroll = "loadMore">
        <div class="book" v-for = "book in books" :key = "book.id">
            <img class="book__img" :src="book.image" alt="">
            <div class="book__detail">
                <h3 class="book__title">{{book.title}}</h3>
                <p class="book__summary">{{book.summary}}</p>
            </div>
        </div>
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
                        "image":"/static/images/1.jpg"
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
