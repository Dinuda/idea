<template>

<div style="overflow: visible;">
    
    <div id="preloader-active" style="display: none;">
        <div class="preloader d-flex align-items-center justify-content-center">
            <div class="preloader-inner position-relative">
                <div class="preloader-circle"></div>
                <div class="preloader-img pere-text">
                    <img src="assets/img/logo/loder.jpg" alt="">
                </div>
            </div>
        </div>
    </div>
    <header>
        <!-- Header Start -->
        <div class="header-area header_area header-transparent">
            <div class="main-header">
                <div class="header-bottom  header-sticky header-color">
                    <div class="container-fluid">
                        <div class="row align-items-center">
                            <!-- Logo -->
                            <div class="col-xl-2 col-lg-1 col-md-1">
                                <div class="logo">
                                    <a href="index.html"><img src="assets/img/logo/logo.png" alt=""></a>
                                </div>
                            </div>
                            <div class="col-xl-8 col-lg-8 col-md-8">
                                <!-- Main-menu -->
                                <div class="main-menu f-right d-none d-lg-block">
                                    <nav>
                                        <ul id="navigation">
                                            <li><a href="index.html">Home</a></li>
                                            <li><a href="about.html">About</a></li>
                                            <li><a href="services.html">Services</a></li>
                                            <li><a href="work.html">Works</a></li>
                                            <li><a href="blog.html">Blog</a>
                                                <ul class="submenu">
                                                    <li><a href="blog.html">Blog</a></li>
                                                    <li><a href="single-blog.html">Blog Details</a></li>
                                                </ul>
                                            </li>
                                            <li><a href="#">Pages</a>
                                                <ul class="submenu">
                                                    <li><a href="elements.html">Element</a></li>
                                                </ul>
                                            </li>
                                            <li><a href="/#/sign_config">Sign In</a></li>
                                        </ul>
                                    </nav>
                                </div>
                            </div>
                            <div class="col-xl-2 col-lg-3 col-md-3">
                                <div class="header-right-btn f-right d-none d-lg-block">
                                    <a href="#" class="btn header-btn">Contact Now</a>
                                </div>
                            </div>
                            <!-- Mobile Menu -->
                            <div class="col-12">
                                <div class="mobile_menu d-block d-lg-none">
                                    <div class="slicknav_menu"><a href="#" aria-haspopup="true" role="button" tabindex="0" class="slicknav_btn slicknav_collapsed" style="outline: none;"><span class="slicknav_menutxt">MENU</span><span class="slicknav_icon"><span class="slicknav_icon-bar"></span><span class="slicknav_icon-bar"></span><span class="slicknav_icon-bar"></span></span></a>
                                        <ul class="slicknav_nav slicknav_hidden" aria-hidden="true" role="menu" style="display: none;">
                                            <li><a href="index.html" role="menuitem" tabindex="-1">Home</a></li>
                                            <li><a href="about.html" role="menuitem" tabindex="-1">About</a></li>
                                            <li><a href="services.html" role="menuitem" tabindex="-1">Services</a></li>
                                            <li><a href="work.html" role="menuitem" tabindex="-1">Works</a></li>
                                            <li class="slicknav_collapsed slicknav_parent"><a href="#" role="menuitem" aria-haspopup="true" tabindex="-1" class="slicknav_item slicknav_row" style="outline: none;"><a href="blog.html" tabindex="-1">Blog</a>
                                                    <span class="slicknav_arrow">+</span></a>
                                                <ul class="submenu slicknav_hidden" role="menu" aria-hidden="true" style="display: none;">
                                                    <li><a href="blog.html" role="menuitem" tabindex="-1">Blog</a></li>
                                                    <li><a href="single-blog.html" role="menuitem" tabindex="-1">Blog Details</a></li>
                                                </ul>
                                            </li>
                                            <li class="slicknav_collapsed slicknav_parent"><a href="#" role="menuitem" aria-haspopup="true" tabindex="-1" class="slicknav_item slicknav_row" style="outline: none;"><a href="#" tabindex="-1">Pages</a>
                                                    <span class="slicknav_arrow">+</span></a>
                                                <ul class="submenu slicknav_hidden" role="menu" aria-hidden="true" style="display: none;">
                                                    <li><a href="elements.html" role="menuitem" tabindex="-1">Element</a></li>
                                                </ul>
                                            </li>
                                            <li><a href="contact.html" role="menuitem" tabindex="-1">Contact</a></li>
                                        </ul>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <!-- Header End -->
    </header>
    <!-- Scroll Up -->
    <div id="back-top" >
        <a title="Go to Top" href="#"> <i class="fas fa-level-up-alt"></i></a>
    </div>

</div>
</template>

<script>
export default {
    name: "site-header",
    components: {}
};
</script>
<style scoped>
.header-color {
    background: #010E21;
}
.header-transparent {
    position: fixed;
    top: 0;
    right: 0;
    left: 0;
    z-index: 9;
}
</style>