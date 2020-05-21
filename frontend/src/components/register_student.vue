<template>
<div class="register container">
    <div class="row">
        <div class="col-12">
            <h2 class="contact-title">Register as Student</h2>
        </div>
        <div class="col-lg-8">
            <div v-if="RegistationUser">
                <div class="form-contact contact_form">
                    <div class="row">
                        <div class="col-12">
                        </div>
                        <div class="col-sm-6">
                            <div class="form-group">
                                <input class="form-control valid" v-model="RegistationUser.username" name="username" type="text" placeholder="Username" required="required">
                            </div>
                        </div>
                        <div class="col-sm-6">
                            <div class="form-group">
                                <input class="form-control valid" v-model="RegistationUser.firstname" name="firstname" type="text" placeholder="First Name" required="required">
                            </div>
                        </div>
                        <div class="col-sm-6">
                            <div class="form-group">
                                <input class="form-control valid" v-model="RegistationUser.lastname" name="lastname" type="text" placeholder="Last Name" required="required">
                            </div>
                        </div>
                        <div class="col-sm-6">
                            <div class="form-group">
                                <select class="form-control valid" required="required">
                                    <option  v-for = " proffesion in professions" :key="proffesion.id">{{ proffesion.name }}</option>
                                </select>
                            </div>
                            <!-- <select>
                                    <option v-for="Proffesion in RegistationUser.proffesions" :key="Proffesion.ID" :value="Proffesion.ID">
                                        {{ Proffesion.Name }}
                                    </option>
                            </select> -->
                        </div>
                        <div class="col-sm-6">
                            <div class="form-group">
                                <input class="form-control valid" v-model="RegistationUser.CV" name="url" type="url" placeholder="Url of you CV" required="required">
                            </div>
                        </div>
                        <div class="col-sm-6">
                            <div class="form-group">
                                <input class="form-control valid" v-model="RegistationUser.email" name="email" type="text" placeholder="E-mail" required="required">
                            </div>
                        </div>
                    </div>
                    <div class="form-group mt-3">
                        <button type="submit" v-on:click="saveUser" class="button button-contactForm boxed-btn">Register</button>
                    </div>
                </div>
            </div>
        </div>
    </div>
</div>
</template>

<script>
import axios from "axios";
export default {
    name: 'register_student',
    data: function () {
        return {
            RegistationUser: {
                username: null,
                firstname: null,
                lastname: null,
                proffesions: [],
                CV: null,
                email: null,
                type: "student",
            },
        };
    },
    components: {

    },
    methods: {
        saveUser() {
            let self = this;
            axios
                .put("http://localhost:8000/addUser", self.RegistationUser)
                .then(function (response) {
                    self.$toasted.success(response.data.message);
                })
                .catch(function (e) {
                    console.log(e);
                });
        },
    },
    mounted() {
        let self = this;
        axios
            .get("http://localhost:8000/professions")
            .then(function (response) {
                self.proffesions = response.data;
                console.log(self.proffesions);
            })
            .catch(function (e) {
                console.log(e);
            });
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
