<template>
    <nav>
        <v-toolbar class="logo" height="100">
            <v-dialog width="400">
                <template v-slot:activator="{ on }">
                    <v-btn v-on="on">Register</v-btn>
                </template>
                <v-card>
                    <v-card-title>
                        <h2>Registration</h2>
                    </v-card-title>
                    <v-form ref="registrationForm">
                        <v-card-text class="px-3">
                            <v-text-field label="Name" v-model="form.name" :rules="rules.nameRules"></v-text-field>
                            <v-text-field label="Surname" v-model="form.surname" :rules="rules.surnameRules"></v-text-field>
                            <v-text-field label="Username" v-model="form.username" :rules="rules.usernameRules"></v-text-field>
                            <v-text-field label="Password" type="password" v-model="form.passowrd" :rules="rules.passwordRules"></v-text-field>
                            <v-text-field label="Repeat password" type="password" v-model="form.repeatedPassword" :rules="rules.passwordRepeatRules"></v-text-field>
                            <v-text-field label="Email" v-model="form.email" :rules="rules.emailRules"></v-text-field>
                            <v-text-field label="Phone number" v-model="form.phoneNumber" :rules="rules.phoneNumberRules"></v-text-field>
                            <label>Choose gender:</label>
                            <v-radio-group v-model="form.gender" required>
                                <v-radio label="Male" value="MALE" checked/>
                                <v-radio label="Female" value="FEMALE"/>
                            </v-radio-group>
                            <v-menu max-width="290">
                                <template v-slot:activator="{ on }">
                                    <v-text-field :value="birthday" placeholder="Birthday date" v-on="on" :rules="rules.dateRules"></v-text-field>
                                </template>
                                <v-date-picker v-model="birthday"></v-date-picker>
                            </v-menu>
                            <label>Choose role:</label>
                            <v-radio-group v-model="form.role" required>
                                <v-radio label="Regular" value="REGULAR" checked/>
                                <v-radio label="Agent" value="AGENT"/>
                            </v-radio-group>
                            <v-text-field label="Web site" v-model="form.website" :rules="rules.websiteRules"></v-text-field>
                            <v-textarea label="Biography" v-model="form.biography" :rules="rules.biographyRules"></v-textarea>
                            <v-btn class="success" @click="submitRegistration">Submit</v-btn>
                        </v-card-text>
                    </v-form>
                </v-card>
            </v-dialog>
            <!--<Registration></Registration>-->
            <!-- <Login></Login> -->

            <v-dialog>
                <template v-slot:activator="{ on }">
                    <v-btn v-on="on">Login</v-btn>
                </template>
                <v-card>
                    <h2>Login</h2>
                    <v-form ref="loginForm">
                        <v-card-text>
                            <v-text-field label="Username" v-model="form.usernameLogin" :rules="rules.usernameRules"></v-text-field>
                            <v-text-field label="Password" type="password" v-model="form.passwordLogin" :rules="rules.passwordRules"></v-text-field>
                            <v-btn class="success" @click="submitLogin">Login</v-btn>
                        </v-card-text>
                    </v-form>
                </v-card>
            </v-dialog> 
            <v-img src="../assets/logoedited2.png" width="20" height="100" position="center"></v-img>
        </v-toolbar>   
    </nav>
</template>

<script>

//import Registration from '../components/Registration'
//import Login from './Login'

export default {
    name: 'LogoToolbar',
    components: {
        //Registration
        //Login,
    },
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
                gender : "",
                birthday : "",
                role : "",
                website: "",
                biography: "",
                usernameLogin: "",
                passwordLogin: ""
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
                     repeatedPassword => !!repeatedPassword || 'Password is required!'
                     //repeatedPassword => this.form.password === repeatedPassword ||  "Passwords do not match!",
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
                ],
                biographyRules: [
                    biography => !!biography || 'Biography can not be empty'
                ]
            },
        }
    },
    methods: {
        submitRegistration(){
            if(this.$refs.registrationForm.validate()){
                console.log(this.form.name, this.form.surname)
            }
        },
         submitLogin(){
             if(this.$refs.loginForm.validate()){
                 console.log(this.form.usernameLogin)
             }
         }
    }
}
</script>