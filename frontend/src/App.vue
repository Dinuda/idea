<template>
<div id="app">
    <noscript>
        <strong>We're sorry but <%= htmlWebpackPlugin.options.title %> doesn't work properly without JavaScript enabled. Please enable it to continue.</strong>
    </noscript>
    <site-header></site-header>
    <router-view></router-view>
    <siteFooter></siteFooter>
</div>
</template>

<script>
import siteHeader from './components/header.vue';
import siteFooter from './components/footer.vue';

export default {
    name: 'App',
    components: {
        siteHeader,
        siteFooter
    }
}
</script>

<style>
#app {
    font-family: Avenir, Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    text-align: center;
    color: #2c3e50;
    margin-top: 60px;
}
</style>
