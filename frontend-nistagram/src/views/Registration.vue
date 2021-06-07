<template>
  <v-app class="grey lighten-4"> 
      <v-container center>
        <v-row>
        <v-col width="300px"></v-col>
        <v-col width="600px">       
          <v-toolbar flat height="45" color="#A29D9C" width="600px">
            <v-app-bar app>
                <v-row>
                    <v-col>
                    <v-toolbar-title >
                        <span outlined class="font-weight-light">NISTA</span>
                        <span>GRAM</span>
                    </v-toolbar-title>
                    </v-col>            
                </v-row>
              </v-app-bar>
          </v-toolbar>
          <v-container>
              <v-form ref="registrationForm">
                <v-row > 
                    <v-col>                       
                            <v-text-field label="Name*" v-model="form.name" :rules="rules.nameRules"></v-text-field>
                            <v-text-field label="Surname*" v-model="form.surname" :rules="rules.surnameRules"></v-text-field>
                            <v-text-field label="Username*" v-model="form.username" :rules="rules.usernameRules"></v-text-field>
                            <v-text-field label="Password*" type="password" v-model="form.password" :rules="rules.passwordRules"></v-text-field>
                            <v-text-field label="Repeat password*" type="password" v-model="form.repeatedPassword" :rules="rules.passwordRepeatRules"></v-text-field>
                            <v-text-field label="Email*" v-model="form.email" :rules="rules.emailRules"></v-text-field>
                            <v-text-field label="Phone number*" v-model="form.phoneNumber" :rules="rules.phoneNumberRules"></v-text-field>
                            <label>Choose gender*</label>
                            <v-radio-group v-model="form.gender" required>
                                <v-radio label="Male" value="MALE"/>
                                <v-radio label="Female" value="FEMALE"/>
                            </v-radio-group>
                            <v-menu max-width="290">
                                <template v-slot:activator="{ on }">
                                    <v-text-field :value="form.birthday" placeholder="Birthday date*" v-on="on" :rules="rules.dateRules"></v-text-field>
                                </template>
                                <v-date-picker v-model="form.birthday"></v-date-picker>
                            </v-menu>
                            <label>Choose role*</label>
                            <v-radio-group v-model="form.role" required>
                                <v-radio label="Regular" value="REGULAR" checked/>
                                <v-radio label="Agent" value="AGENT"/>
                            </v-radio-group>
                            <v-text-field label="Web site*" v-model="form.website" :rules="rules.websiteRules"></v-text-field>
                            <v-textarea label="Biography" v-model="form.biography" :rules="rules.biographyRules"></v-textarea>
                            <v-btn class="success" @click="submit">Submit</v-btn>
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

export default {
    name: 'Registration',
    data() {
        return {
            form: {
                name: "",
                surname: "",
                username: "",
                password: "",
                repeatedPassword: "",
                email: "",
                phoneNumber: "",
                gender : 1,
                birthday : "",
                role : 1,
                website: "",
                biography: "",
            },
            rules:{
                emailRules : [
                    email => !!email || 'Email is required!',
                    email =>  /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/.test(email) || 'Email must be valid!'
                ],
                passwordRules: [
                    password => !!password || 'Password is required!',
                ],
                 passwordRepeatRules: [
                     repeatedPassword => !!repeatedPassword || 'Password is required!',
                     repeatedPassword => this.form.password === repeatedPassword ||  "Passwords do not match!",
                ],
                nameRules: [
                        name => !!name || 'Name is required!',
                        name => /^[A-Za-z\s]+$/.test(name) || 'Name must contain only letters!'
                ],
                surnameRules: [
                        surname => !!surname || 'Surname is required!',
                        surname => /^[A-Za-z\s]+$/.test(surname) || 'Surname must contain only letters!'
                ],
                phoneNumberRules: [
                        phoneNumber => !!phoneNumber || 'Phone number is required!',
                        phoneNumber => (phoneNumber && phoneNumber.length >=9) || 'Phone number imust have at least 9 digits!',
                        phoneNumber => /^[0-9\s]+$/.test(phoneNumber) || 'Phone number must contain only numbers!'
                ],
                usernameRules: [
                    username => !!username || 'Username is rquired',
                    username => /^[0-9A-Za-z\s]+$/.test(username) || 'Username can only contain numbers and letters!'
                ],
                websiteRules: [
                    website => !!website || 'Website is required'
                    //website => /^((ftp|http|https):\/\/)?(www.)?(?!.*(ftp|http|https|www.))[a-zA-Z0-9_-]+(\.[a-zA-Z]+)+((\/)[\w#]+)*(\/\w+\?[a-zA-Z0-9_]+=\w+(&[a-zA-Z0-9_]+=\w+)*)?$/.test(website) || 'Website not valid!'
                ],
                dateRules: [
                    birthday => !!birthday || 'Birthday is required',
                    birthday => /^((?:19|20)[0-9][0-9])-(0?[1-9]|1[012])-(0?[1-9]|[12][0-9]|3[01])$/.test(birthday) || 'Invalid input'
                ]              
            },
        }
    },
    methods: {
        submit(){
            let RegularUserDTO = {
                Name: this.form.name,
                Surname: this.form.surname,
                Username: this.form.username,
                Password: this.form.password,
                Email: this.form.email,
                PhoneNumber: this.form.phoneNumber,
                Gender: 1,
                BirthDate: this.form.birthday,
                UserRole: this.form.userRole,
                UserType: 1,
                Biography: this.form.biography,
                WebSite: this.form.website,
                IsDisabled: false,
                PrivacyType: 1,
                AllMessageRequests: true,
                TagsAllowed: true,               
            }
            axios.post('http://localhost:8081/create-regular-user', RegularUserDTO)
            .then(response=>{
                console.log(response.status);
                
                setTimeout(() => { this.$router.push({path: '/'})}, 2000);
            })
            
        },
    }
}
</script>

<style>

</style>