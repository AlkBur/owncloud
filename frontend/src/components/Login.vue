<template>
    <!-- Main Content -->
    <div class="container-fluid">
        <div class="row main-content bg-success text-center">
            <div class="col-md-4 text-center company__info">
                <span class="company__logo"
                    ><h1>
                        <font-awesome-icon icon="user" /></h1
                ></span>
                <h4 class="company_title">Own cloud</h4>
            </div>
            <div class="col-md-8 col-xs-12 col-sm-12 login_form">
                <div class="container-fluid">
                    <div class="row">
                        <h2>Log In</h2>
                    </div>
                    <div class="row">
                        <Form
                            @submit="handleLogin"
                            :validation-schema="schema"
                            class="form-group"
                        >
                            <div class="form-group">
                                <Field
                                    type="text"
                                    name="username"
                                    id="username"
                                    class="form__input"
                                    placeholder="Username"
                                />
                                <ErrorMessage
                                    name="username"
                                    class="error-feedback"
                                />
                            </div>
                            <div class="form-group">
                                <!-- <span class="fa fa-lock"></span> -->
                                <Field
                                    type="password"
                                    name="password"
                                    id="password"
                                    class="form__input"
                                    placeholder="Password"
                                />
                                <ErrorMessage
                                    name="password"
                                    class="error-feedback"
                                />
                            </div>
                            <div class="row">
                                <input
                                    type="submit"
                                    value="Submit"
                                    class="btn"
                                />
                            </div>
                        </Form>
                    </div>
                </div>
            </div>
        </div>
    </div>

    <!-- Footer -->
    <div
        v-if="message"
        class="container-fluid text-center footer alert alert-danger"
        role="alert"
    >
        {{ message }}
    </div>

    <!-- 





    <div class="col-md-12">
        <div class="card card-container">
            <img
                id="profile-img"
                src="//ssl.gstatic.com/accounts/ui/avatar_2x.png"
                class="profile-img-card"
            />
            <Form @submit="handleLogin" :validation-schema="schema">
                <div class="form-group">
                    <label for="username">Username</label>
                    <Field name="username" type="text" class="form-control" />
                    <ErrorMessage name="username" class="error-feedback" />
                </div>
                <div class="form-group">
                    <label for="password">Password</label>
                    <Field
                        name="password"
                        type="password"
                        class="form-control"
                    />
                    <ErrorMessage name="password" class="error-feedback" />
                </div>

                <div class="form-group">
                    <button
                        class="btn btn-primary btn-block"
                        :disabled="loading"
                    >
                        <span
                            v-show="loading"
                            class="spinner-border spinner-border-sm"
                        ></span>
                        <span>Login</span>
                    </button>
                </div>

                <div class="form-group">
                    <div v-if="message" class="alert alert-danger" role="alert">
                        {{ message }}
                    </div>
                </div>
            </Form>
        </div>
    </div> -->
</template>

<script>
import { Form, Field, ErrorMessage } from "vee-validate";
import * as yup from "yup";

export default {
    name: "Login",
    components: {
        Form,
        Field,
        ErrorMessage,
    },
    data() {
        const schema = yup.object().shape({
            username: yup.string().required("Username is required!"),
            password: yup.string().required("Password is required!"),
        });

        return {
            loading: false,
            message: "",
            schema,
        };
    },
    computed: {
        loggedIn() {
            return this.$store.state.auth.status.loggedIn;
        },
    },
    created() {
        if (this.loggedIn) {
            this.$router.push("/profile");
        }
    },
    methods: {
        handleLogin(user) {
            this.loading = true;

            this.$store.dispatch("auth/login", user).then(
                () => {
                    this.$router.push("/profile");
                },
                (error) => {
                    this.loading = false;
                    this.message =
                        (error.response &&
                            error.response.data &&
                            error.response.data.message) ||
                        error.message ||
                        error.toString();
                }
            );
        },
    },
};
</script>

<style scoped>
.main-content {
    width: 50%;
    border-radius: 20px;
    box-shadow: 0 5px 5px rgba(0, 0, 0, 0.4);
    margin: 5em auto;
    display: flex;
}
.company__info {
    background-color: #008080;
    border-top-left-radius: 20px;
    border-bottom-left-radius: 20px;
    display: flex;
    flex-direction: column;
    justify-content: center;
    color: #fff;
}
.fa-android {
    font-size: 3em;
}
@media screen and (max-width: 640px) {
    .main-content {
        width: 90%;
    }
    .company__info {
        display: none;
    }
    .login_form {
        border-top-left-radius: 20px;
        border-bottom-left-radius: 20px;
    }
}
@media screen and (min-width: 642px) and (max-width: 800px) {
    .main-content {
        width: 70%;
    }
}
.row > h2 {
    color: #008080;
}
.login_form {
    background-color: #fff;
    border-top-right-radius: 20px;
    border-bottom-right-radius: 20px;
    border-top: 1px solid #ccc;
    border-right: 1px solid #ccc;
}
form {
    padding: 0 2em;
}
.form__input {
    width: 100%;
    border: 0px solid transparent;
    border-radius: 0;
    border-bottom: 1px solid #aaa;
    padding: 1em 0.5em 0.5em;
    padding-left: 2em;
    outline: none;
    margin: 1.5em auto;
    transition: all 0.5s ease;
}
.form__input:focus {
    border-bottom-color: #008080;
    box-shadow: 0 0 5px rgba(0, 80, 80, 0.4);
    border-radius: 4px;
}
.btn {
    transition: all 0.5s ease;
    width: 70%;
    border-radius: 30px;
    color: #008080;
    font-weight: 600;
    background-color: #fff;
    border: 1px solid #008080;
    margin-top: 1.5em;
    margin-bottom: 1em;
}
.btn:hover,
.btn:focus {
    background-color: #008080;
    color: #fff;
}
</style>
