<template>
<div class="slider-area  slider-height background" data-background="assets/img/hero/h1_hero.jpg" style="background-image: url(&quot;assets/img/hero/h1_hero.jpg&quot;);">
    <div class="slider-active slick-initialized slick-slider">
        <!-- Single Slider -->
        <div class="slick-list draggable">
            <div class="slick-track" style="opacity: 1; width: 3226px;">
                <div class="single-slider slick-slide slick-current slick-active" data-slick-index="0" aria-hidden="false" tabindex="0" style="width: 1613px; position: relative; left: 0px; top: 0px; z-index: 999; opacity: 1;">
                    <div class="slider-cap-wrapper">
                        <div class="hero__img">
                            <img src="../assets/images/hero.jpg" alt="" data-animation="fadeInRight" data-transition-duration="5s" class="">
                        </div>
                        <div class="hero__caption">
                            <h1 data-animation="fadeInLeft" data-delay=".4s" class="" style="animation-delay: 0.4s;">Business Synergy via <br>Robust Strategies</h1>
                            <p data-animation="fadeInLeft" data-delay=".6s" class="" style="animation-delay: 0.6s;">Aasrem dfgum dolor sivst amet, consectetur adipisicing elitvde, sed dvo eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad micnim veniam, quis nostrud exercita.</p>
                            <!-- Hero Btn -->
                            <a href="#" class="hero-btn" data-animation="fadeInLeft" data-delay=".9s" tabindex="0" style="animation-delay: 0.9s;">
                                View project
                                <!-- Arrow SVG icon -->
                            </a>
                        </div>
                    </div>
                </div>
                <div class="single-slider slick-slide" data-slick-index="1" aria-hidden="true" tabindex="-1" style="width: 1613px; position: relative; left: -1613px; top: 0px; z-index: 998; opacity: 0;">
                    <div class="slider-cap-wrapper">
                        <div class="hero__img">
                            <img src="assets/img/hero/hero_img2.jpg" alt="" data-animation="fadeInRight" data-transition-duration="5s">
                        </div>
                        <div class="hero__caption">
                            <a href="#" class="hero-btn" data-animation="fadeInLeft" data-delay=".9s" tabindex="-1">
                                View project
                                <!-- Arrow SVG icon -->
                                <svg version="1.1" id="Capa_1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" x="0px" y="0px" viewBox="0 0 512 512" style="enable-background:new 0 0 512 512;" xml:space="preserve">
                                    <path d="M506.134,241.843c-0.006-0.006-0.011-0.013-0.018-0.019l-104.504-104c-7.829-7.791-20.492-7.762-28.285,0.068
                                         c-7.792,7.829-7.762,20.492,0.067,28.284L443.558,236H20c-11.046,0-20,8.954-20,20c0,11.046,8.954,20,20,20h423.557
                                         l-70.162,69.824c-7.829,7.792-7.859,20.455-0.067,28.284c7.793,7.831,20.457,7.858,28.285,0.068l104.504-104
                                         c0.006-0.006,0.011-0.013,0.018-0.019C513.968,262.339,513.943,249.635,506.134,241.843z"></path>
                                </svg>
                            </a>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</div>
</template>

<script>
export default {
    name: 'hero',
    components: {}
}
</script>
<style scoped>
.background {
    background-color: #010E21;
}
</style>
