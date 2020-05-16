<template>
<div class="register container">
    <div class="counter_form ">
        <div class="row fill_height">
            <div class="col fill_height">
                <form class="counter_form_content d-flex flex-column align-items-center justify-content-center" action='/#/submit_user_signup' method="POST" >
                    <div class="counter_form_title">Register to buy courses</div>
                    <input type="tel" name="username" class="counter_input" placeholder="UserName:" required="required">
                    <input type="tel" name="password" class="counter_input" placeholder="Password:" required="required">
                    <input type="text" name="firstnamename" class="counter_input" placeholder="Your Name:" required="required">
                    <input type="text" name="lastnamename" class="counter_input" placeholder="Your Name:" required="required">
                    <input type="tel" name="phone" class="counter_input" placeholder="Phone:">
                    <input type="tel" name="email" class="counter_input" placeholder="E-mail:" required="required">
                    <input type="tel" name="child_name" class="counter_input" placeholder="Child's Name:" required="required">
                    <select name="counter_select" id="counter_select"  class="counter_input counter_options" placeholder="Your interests">

                        <option id="{professionID}">Your Interests</option>
                        <option>Programming</option>
                        <option>Creating</option>
                        <option>Problem solving</option>
                    </select>
                    <button type="submit" class="counter_form_button">Register</button>
                </form>
            </div>
        </div>
    </div>
</div>
</template>

<script>
export default {
    name: 'register',
    components: {

    },
    methods: {
        fetch()
    }
}
</script>

<style scoped>
.register {
  padding-top: 150px;
}
</style>