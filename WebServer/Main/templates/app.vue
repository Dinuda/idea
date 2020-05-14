<template>
<div id="app" class="super_container">
    <noscript>
        <strong>We're sorry but <%= htmlWebpackPlugin.options.title %> doesn't work properly without JavaScript enabled. Please enable it to continue.</strong>
    </noscript>
    <site-header></site-header>
    <br><br><br><br>
    <router-view></router-view>
    <site-footer></site-footer>

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
@import "/styles/main_styles.css";
@import "/styles/responsive.css";

@import "/styles/courses.css";
@import "/styles/courses_responsive.css";

@import "/styles/course.css";
@import "/styles/course_responsive.css";

@import "./assets/css/style.css";
</style>
