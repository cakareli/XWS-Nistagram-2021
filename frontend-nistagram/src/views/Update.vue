<template>
  <v-app class="grey lighten-4">
    <v-container center>
      <v-row>
        <v-col width="300px"></v-col>
        <v-col width="600px">
          <v-toolbar height="45" color="#A29D9C" width="600px">
            <v-app-bar app height="45">
              <v-app-bar-nav-icon @click="$router.push('/account')">
                <v-icon>mdi-arrow-left</v-icon>
              </v-app-bar-nav-icon>
              <v-row>
                <v-col width="550px">
                  <v-toolbar-title>
                    <span outlined class="font-weight-light">NISTA</span>
                    <span>GRAM</span>
                  </v-toolbar-title>
                </v-col>
              </v-row>
            </v-app-bar>
          </v-toolbar>
          <v-container>
            <v-snackbar v-model="snackbar" top timeout="3500">
              <span>{{ snackbarText }}</span>
            </v-snackbar>
            <v-form ref="updateForm">
              <v-row>
                <v-col>
                  <v-text-field
                    label="Name*"
                    v-model="form.name"
                    :rules="rules.nameRules"
                  ></v-text-field>
                  <v-text-field
                    label="Surname*"
                    v-model="form.surname"
                    :rules="rules.surnameRules"
                  ></v-text-field>
                  <v-text-field
                    label="Username*"
                    v-model="form.username"
                    :rules="rules.usernameRules"
                  ></v-text-field>
                  <v-text-field
                    label="Email*"
                    v-model="form.email"
                    :rules="rules.emailRules"
                  ></v-text-field>
                  <v-text-field
                    label="Phone number*"
                    v-model="form.phoneNumber"
                    :rules="rules.phoneNumberRules"
                  ></v-text-field>
                  <label>Choose gender*</label>
                  <v-radio-group v-model="form.gender" required>
                    <v-radio label="Male" value="0" />
                    <v-radio label="Female" value="1" />
                  </v-radio-group>
                  <v-menu max-width="290">
                    <template v-slot:activator="{ on }">
                      <v-text-field
                        :value="form.birthday"
                        v-model="form.birthday"
                        placeholder="Birthday date*"
                        v-on="on"
                        :rules="rules.dateRules"
                      ></v-text-field>
                    </template>
                    <v-date-picker v-model="form.birthday"></v-date-picker>
                  </v-menu>
                  <v-text-field
                    label="Web site*"
                    v-model="form.website"
                    :rules="rules.websiteRules"
                  ></v-text-field>
                  <v-textarea
                    label="Biography"
                    v-model="form.biography"
                    :rules="rules.biographyRules"
                  ></v-textarea>
                  <v-btn class="success" @click="update">Apply changes</v-btn>
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
import axios from "axios";
import { getId, getToken } from '../security/token.js'

export default {
  name: "Update",

  mounted() {
      this.loadRegisteredUser();
  },

  data() {
    return {
      snackbar: false,

      snackbarText: "",

      form: {
        name: "",
        surname: "",
        username: "",
        email: "",
        phoneNumber: "",
        gender: "",
        birthday: "",
        website: "",
        biography: "",
      },

      time: "T15:00:00+01:00",

      rules: {
        emailRules: [
          (email) => !!email || "Email is required!",
          (email) =>
            /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/.test(
              email
            ) || "Email must be valid!",
        ],
        passwordRules: [(password) => !!password || "Password is required!"],
        passwordRepeatRules: [
          (repeatedPassword) => !!repeatedPassword || "Password is required!",
          (repeatedPassword) =>
            this.form.password === repeatedPassword ||
            "Passwords do not match!",
        ],
        nameRules: [
          (name) => !!name || "Name is required!",
          (name) =>
            /^[A-Za-z\s]+$/.test(name) || "Name must contain only letters!",
        ],
        surnameRules: [
          (surname) => !!surname || "Surname is required!",
          (surname) =>
            /^[A-Za-z\s]+$/.test(surname) ||
            "Surname must contain only letters!",
        ],
        phoneNumberRules: [
          (phoneNumber) => !!phoneNumber || "Phone number is required!",
          (phoneNumber) =>
            (phoneNumber && phoneNumber.length >= 9) ||
            "Phone number imust have at least 9 digits!",
          (phoneNumber) =>
            /^[0-9\s]+$/.test(phoneNumber) ||
            "Phone number must contain only numbers!",
        ],
        usernameRules: [
          (username) => !!username || "Username is rquired",
          (username) =>
            /^[0-9A-Za-z\s]+$/.test(username) ||
            "Username can only contain numbers and letters!",
        ],
        websiteRules: [
          (website) => !!website || "Website is required",
          //website => /^((ftp|http|https):\/\/)?(www.)?(?!.*(ftp|http|https|www.))[a-zA-Z0-9_-]+(\.[a-zA-Z]+)+((\/)[\w#]+)*(\/\w+\?[a-zA-Z0-9_]+=\w+(&[a-zA-Z0-9_]+=\w+)*)?$/.test(website) || 'Website not valid!'
        ],
        dateRules: [
          (birthday) => !!birthday || "Birthday is required",
          (birthday) =>
            /^((?:19|20)[0-9][0-9])-(0?[1-9]|1[012])-(0?[1-9]|[12][0-9]|3[01])$/.test(
              birthday
            ) || "Invalid input",
        ],
      },
    };
  },
  methods: {
    update() {
      if (this.$refs.updateForm.validate()) {
        axios
          .post("http://localhost:8081/api/user/update-regular-user", {
            _id: getId(),
            name: this.form.name,
            surname: this.form.surname,
            username: this.form.username,
            email: this.form.email,
            phoneNumber: this.form.phoneNumber,
            gender: parseInt(this.form.gender, 10),
            birthDate: this.form.birthday + this.time,
            biography: this.form.biography,
            webSite: this.form.website,
          })
          .then((response) => {
            console.log(response.status);
            this.snackbarText = "Your account has been successfuly updated!";
            this.snackbar = true;
            setTimeout(() => {
              this.$router.push({ path: "/account" });
            }, 2000);
          })
          .catch((error) => {
            if (error.response.status === 400) {
              this.snackbarText = "Bad request!";
              this.snackbar = true;
            } else if (error.response.status === 409) {
              this.snackbarText =
                "User with same email address already exists!";
              this.snackbar = true;
            }
          });
      } else {
        this.snackbarText = "Please fill all fields correctly!";
        this.snackbar = true;
      }
    },

    loadRegisteredUser(){
            axios.get('http://localhost:8081/api/user/logged-user/'+getId(),
            {
                    headers : {
                        Authorization: 'Bearer ' + getToken()
                    }
            })
            .then(response => {
                this.regularUser = response.data;
                this.form.id = this.regularUser.id;
                this.form.name = this.regularUser.Name;
                this.form.surname = this.regularUser.Surname;
                this.form.username = this.regularUser.Username;
                this.form.gender = String(this.regularUser.Gender);
                this.form.phoneNumber = this.regularUser.PhoneNumber;
                this.form.email = this.regularUser.Email;
                this.form.birthday = this.regularUser.BirthDate.substring(0,10);
                this.form.biography = this.regularUser.Biography;
                this.form.website = this.regularUser.WebSite;
            }).catch(error => {
            if(error.response.status === 500){
                this.snackbarText = "Internal server error occurred!";
                this.snackbar = true;
            }
            if(error.response.status === 401){
               this.snackbar = true
               this.snackbarText = "You are unauthorized to get patient informations!";
            }
        })
    },
  },
};
</script>

