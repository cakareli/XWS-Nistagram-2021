<template>
  <v-app class="grey lighten-4">
    <v-container center>
      <v-row>
        <v-col width="300px"></v-col>
        <v-col width="600px">
          <v-toolbar height="45" width="800px">
            <v-app-bar app height="45" color="grey">
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
            <v-form ref="verificationRequestForm">
              <v-row>
                <v-col>
                  <v-text-field
                    label="Real Name*"
                    v-model="form.name"
                    :rules="rules.nameRules"
                  ></v-text-field>
                  <v-text-field
                    label="Real Surname*"
                    v-model="form.surname"
                    :rules="rules.surnameRules"
                  ></v-text-field>
                  <label>Upload image of your official document (ID, drivers licence, passport)</label>
                  <br/>
                  <br/>
                  <v-btn raised class="grey lighten-3" @click="uploadFile"
                    >Upload Image</v-btn
                  >
                  <input
                    type="file"
                    style="display: none"
                    ref="fileInput"
                    accept="image/*"
                    @change="onFilePicked"
                  />
                  <br />
                  <br />
                  <v-col>
                        <img :src="imageUrl" height="500" width="400"/>
                  </v-col>
                  <br />
                  <br />
                  <label>Choose user type*</label>
                  <v-radio-group v-model="form.verificationType" required>
                    <v-radio label="Influencer" value="1" />
                    <v-radio label="Sports" value="2" />
                    <v-radio label="Business" value="3" />
                    <v-radio label="Brand" value="4" />
                    <v-radio label="Organization" value="5" />
                    <v-radio label="NewMedia" value="6" />
                  </v-radio-group>
                  <v-btn class="success" @click="sendVerificationRequest"
                    >Send verification request</v-btn
                  >
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
import firebase from 'firebase/app'
import 'firebase/firebase-storage'
import axios from "axios";
import { getId, getToken } from "../security/token.js";

export default {
  name: "VerificationRequest",

  mounted() {
    this.loadRegisteredUser();
  },

  data() {
    return {
      imageUrl: "",
      image: null,
      uploadValue: 0,
      firebaseURL: [],
      snackbar: false,
      snackbarText: "",
      form: {
        name: "",
        surname: "",
        imageUrl: "",
        verificationType: "1",
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
    uploadFile() {
      this.$refs.fileInput.click();
    },
    onFilePicked(event) {
      const files = event.target.files;
      let filename = files[0].name;
      console.log(filename);
      if (filename.lastIndexOf(".") <= 0) {
        return alert("Please upload valid image file!");
      }

      const fileReader = new FileReader();
      fileReader.addEventListener("load", () => {
        this.imageUrl = fileReader.result;
      });
      fileReader.readAsDataURL(files[0]);
      this.image = files[0];

      const storageRef = firebase.storage().ref(`${filename}`).put(this.image);
      storageRef.on(
        `state_changed`,
        (snapshot) => {
          this.uploadValue =
            (snapshot.bytesTransferred / snapshot.totalBytes) * 100;
        },
        (error) => {
          console.log(error.message);
        },
        () => {
          this.uploadValue = 100;
          storageRef.snapshot.ref.getDownloadURL().then((url) => {
            this.firebaseURL.push(url);
            console.log(this.firebaseURL);
          });
        }
      );
    },
    sendVerificationRequest() {
      if (this.$refs.verificationRequestForm.validate()) {
        axios
          .post("http://localhost:8081/api/user/verification-request", {
            _id: getId(),
            name: this.form.name,
            surname: this.form.surname,
            imageUrl: this.firebaseURL[0],
            verificationType: parseInt(this.form.verificationType, 10),
          })
          .then((response) => {
            console.log(response.status);
            this.snackbarText =
              "Your verification request has been successfuly sent!";
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

    loadRegisteredUser() {
      axios
        .get("http://localhost:8081/api/user/logged-user/" + getId(), {
          headers: {
            Authorization: "Bearer " + getToken(),
          },
        })
        .then((response) => {
          this.regularUser = response.data;
          this.form.id = this.regularUser.id;
          this.form.name = this.regularUser.Name;
          this.form.surname = this.regularUser.Surname;
        })
        .catch((error) => {
          if (error.response.status === 404) {
            this.snackbarText = "Internal server error occurred!";
            this.snackbar = true;
          }
          if (error.response.status === 401) {
            this.snackbar = true;
            this.snackbarText =
              "You are unauthorized to get patient informations!";
          }
        });
    },
  },
};
</script>

