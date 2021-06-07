<template>
    <nav>
        <v-toolbar class="grey">

        </v-toolbar>
        <h1>AAA</h1>
    </nav>
</template>

<script>
export default {
    name: 'UpperToolbar',
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

<style>

</style>