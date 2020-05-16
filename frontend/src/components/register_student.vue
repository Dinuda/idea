<template>
<div class="register container">
    <div class="row">
        <div class="col-12">
            <h2 class="contact-title">Register as Student</h2>
        </div>
        <div class="col-lg-8">
            <form class="form-contact contact_form" action="contact_process.php" method="post" id="contactForm" novalidate="novalidate">
                <div class="row">
                    <div class="col-12">
                    </div>
                    <div class="col-sm-6">
                        <div class="form-group">
                            <input class="form-control valid" name="username" type="text"  placeholder="Username" required="required">
                        </div>
                    </div>
                    <div class="col-sm-6">
                        <div class="form-group">
                            <input class="form-control valid" name="firstname" type="text" placeholder="First Name" required="required">
                        </div>
                    </div>
                     <div class="col-sm-6">
                        <div class="form-group">
                            <input class="form-control valid" name="lastname" type="text" placeholder="Last Name" required="required">
                        </div>
                    </div> <div class="col-sm-6">
                        <div class="form-group">
                            <input class="form-control valid" name="occupation" type="number" placeholder="Occupation" required="required">
                        </div>
                    </div>
                    <div class="col-sm-6">
                        <div class="form-group">
                            <input class="form-control valid" name="url" type="url" placeholder="Url of you CV" required="required">
                        </div>
                    </div>
                    <div class="col-sm-6">
                        <div class="form-group">
                            <input class="form-control valid" name="email" type="text" placeholder="E-mail" required="required">
                        </div>
                    </div>
                    <div class="col-12">
                        <div class="form-group">
                            <input class="form-control valid" name="occupation" type="text" placeholder="Proffesions" required="required">
                        </div>
                    </div>
                </div>
                <div class="form-group mt-3">
                    <button type="submit" class="button button-contactForm boxed-btn">Register</button>
                </div>
            </form>
        </div>

    </div>
</div>
</template>

<script>
export default {
    name: 'register_student',
    components: {

    },
    methods: {
        signup: function () {

        }
    }
}
</script>

<style scoped>
.register {
    padding-top: 150px;
}

.counter_form[data-v-8f8a1d9a] {
    right: 30px;
    width: 380px;
    height: 100%;
    background: #FFFFFF;
    padding-left: 40px;
    padding-right: 40px;
    box-shadow: 0px 5px 40px rgba(29, 34, 47, 0.15);
    margin-bottom: 23px;
    padding-top: 41px;
    padding-bottom: 35px;
    width: 100%;
}
.header-color[data-v-29e8c3c6] {
    background: #010E21;
}

.contact-title {
    margin-right: 1004px;
}
.register[data-v-8f8a1d9a] {
    padding-top: 80px;
}
.boxed-btn[data-v-8f8a1d9a] {
    margin-left: 593px;
}
</style>
