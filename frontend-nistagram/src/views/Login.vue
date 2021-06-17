<template>
  <v-app class="grey lighten-4">
    <v-snackbar v-model="snackbar" centered timeout="3500">
      <span>{{snackbarText}}</span>
      </v-snackbar>  
    <v-container center>
      <v-row>
        <v-col width="300px"></v-col>
        <v-col width="600px">
          <v-toolbar flat height="45" width="800px">
            <v-app-bar app height="45" color="grey lighten-3">
              <v-app-bar-nav-icon @click="$router.push('/')">
                <v-icon>mdi-arrow-left</v-icon>
              </v-app-bar-nav-icon>
              <v-row>
                <v-col>
                  <v-toolbar-title>
                    <span outlined class="font-weight-light">NISTA</span>
                    <span>GRAM</span>
                  </v-toolbar-title>
                </v-col>
              </v-row>
            </v-app-bar>
          </v-toolbar>
          <v-container>
            <v-form ref="loginForm">
              <v-row>
                <v-col>
                  <v-text-field
                    v-model="form.username"
                    label="Username"
                    width="100px"
                    :rules="rules.usernameRules"
                  ></v-text-field>
                  <v-text-field
                    v-model="form.password"
                    type="password"
                    label="Password"
                    width="100px"
                    :rules="rules.passwordRules"
                  ></v-text-field>
                  <v-btn class="success" @click="submit">Login</v-btn>
                </v-col>
              </v-row>
            </v-form>
          </v-container>
        </v-col>
        <v-col width="300px"></v-col>
      </v-row>
    </v-container>
  </v-app>
</template>

<script>
import axios from 'axios';
import { setToken } from '../security/token.js'

export default {
  name: "Login",
  data() {
    return {
      form: {
        username: "",
        password: "",
      },
      rules: {
        passwordRules: [(password) => !!password || "Password is required!"],
        usernameRules: [
          (username) => !!username || "Username is rquired",
          (username) =>
            /^[0-9A-Za-z\s]+$/.test(username) ||
            "Username can only contain numbers and letters!",
        ],
      },
      snackbar: false,
      snackbarText : "",
    };
  },
  methods: {
    submit() {
            let loginCredentials = {
                username: this.form.username,
                password: this.form.password
            }
            axios.post('http://localhost:8081/api/auth/login', loginCredentials)
            .then(response =>{
                console.log(response)
                let token = response.data.token;
                setToken(token);
                this.$router.push({ path: "/" });
            }).catch(error => {
                if(error.response.status === 400){
                    this.snackbarText = "Account with that username doesn't exist!";
                    this.snackbar = true;
                }else if(error.response.status === 500){
                    this.snackbarText = "Internal server error occurred!";
                    this.snackbar = true;
                }else if(error.response.status === 401){
                    this.snackbarText = "You have entered wrong password!";
                    this.snackbar = true;
                }
            })
    },
  },
};
</script>

<style>
</style>