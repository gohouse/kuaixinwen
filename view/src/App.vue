<template>
  <div id="app">
      <div class="header">
          <span class="topbar">快新闻</span>
      </div>
    <router-view/>
  </div>
</template>

<script>
export default {
  name: 'App'
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: left;
  color: #2c3e50;
  margin: 0 auto;
  height:100%;
  width: 90%;
}
.header {
    width: 100%;
    height: 2rem;
    padding:0.5rem 1rem;
    text-align: center;
}
.topbar {
    line-height:2rem;
    font-size:1.5rem;
}


/*@media (min-width: 768px){*/
  /*#app {*/
    /*max-width: 960px;*/
  /*}*/
/*}*/

/*@media (min-width: 992px){ //>=992的设备 }*/

/*@media (min-width: 1200){ //>=1200的设备 }*/
</style>
