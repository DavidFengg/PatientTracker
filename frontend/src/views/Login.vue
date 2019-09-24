<template>
    <div>
        <h1>Login</h1>
        <b-form @submit.prevent="login">
            <b-form-group
                label="Username"
                >
                <b-form-input v-model="username" placeholder=""></b-form-input>
            </b-form-group>

            <b-form-group
                label="Password"
                >
                <b-form-input v-model="password" placeholder=""></b-form-input>
            </b-form-group>
            <b-button type="submit">Submit</b-button>
        </b-form>
    </div>
</template>

<script>
export default {
    data() {
        return {
            username: "",
            password: ""
        };
    },

    methods: {
        login() {
            let username = this.username;
            let password = this.password;

            this.$store.dispatch('login', {username, password})
            .then(() => this.$router.push('/'))
            .catch(err => console.log(err));
        }
    }
}
</script>