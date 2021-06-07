<template>
  <v-app class="grey lighten-4"> 
      <v-container center>
        <v-row>
        <v-col width="300px"></v-col>
        <v-col width="600px">       
          <v-toolbar flat app height="45" color="#A29D9C" width="600px">
            <v-row>
                <v-col>
                  <v-toolbar-title >
                    <span outlined class="font-weight-light">NISTA</span>
                    <span>GRAM</span>
                  </v-toolbar-title>
                </v-col>            
              </v-row>
          </v-toolbar>
          <v-container>
              <v-form ref="loginForm">
                <v-row > 
                    <v-col>
                        <v-text-field v-model="form.username" label="Username" width="100px" :rules="rules.usernameRules"></v-text-field>
                        <v-text-field v-model="form.password" type="password" label="Password" width="100px" :rules="rules.passwordRules"></v-text-field>
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
export default {
    name: 'Login',
    data(){
        return{
            form: {
                username: "",
                password: ""
            },
            rules:{
                passwordRules: [
                            password => !!password || 'Password is required!',
                ],
                usernameRules: [
                            username => !!username || 'Username is rquired',
                            username => /^[0-9A-Za-z\s]+$/.test(username) || 'Username can only contain numbers and letters!'
                ],
            }   
        }
    },
    methods:{
        submit(){
            if(this.$refs.loginForm.validate()){
                 console.log(this.form.username)
        }
    }
    }
}
</script>

<style>

</style>