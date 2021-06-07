<template>
<v-app>
  <v-dialog>
    <template v-slot:activator="{ on }">
        <v-btn v-on="on">Login</v-btn>
    </template>
    <v-card>
        <h2>Login</h2>
    
        <v-form ref="loginForm">
            <v-card-text>
                <v-text-field label="Username" v-model="usernameLogin" :rules="rules.usernameRules"></v-text-field>
                <v-text-field label="Password" type="password" v-model="passowordLogin" :rules="rules.passwordRules"></v-text-field>
                <v-btn class="success" @click="submitLogin">Login</v-btn>
            </v-card-text>
        </v-form>
    </v-card>
</v-dialog>
</v-app>
</template>

<script>
export default {
    data(){
        return{
            usernameLogin: "",
            passowordLogin: ""
        }
    }
}
</script>

<style>

</style>