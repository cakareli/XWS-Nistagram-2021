<template>
    <v-app class="grey lighten-3" width="1000px">
        <v-app-bar app height="45" class="grey lighten-3">
        <v-row>
          <v-col>
            <v-toolbar-title>
              <span outlined class="font-weight-light">NISTA</span>
              <span>GRAM</span>
            </v-toolbar-title>
          </v-col>
          <v-spacer></v-spacer>
          <v-col>
            <v-btn outlined class="mx-5 white">
              <v-icon>mdi-send</v-icon>
            </v-btn>
          </v-col>
        </v-row>
        </v-app-bar>
                <v-container class="justify-center">
                    <v-snackbar v-model="snackbar" top timeout="3500">
                        <span>{{snackbarText}}</span>
                    </v-snackbar>
                    <v-container class="pa-12 justify-center" width="1000px" height="1000">
                    <v-row>
                    <v-col>
                        <v-form v-model="form.isFormValid">
                            <v-btn raised class="grey lighten-3" @click="uploadFile">Upload Image</v-btn>
                            <input type="file" style="display: none" ref="fileInput" accept="image/*" @change="onFilePicked"/>
                            <v-text-field label ="Description" v-model="form.description" :rules="rules.description"></v-text-field>
                            <v-text-field label ="Location" v-model="form.location" :rules="rules.location"></v-text-field>
                            <v-row>
                                <v-text-field label ="Tags" v-model="tag" @click="addTag"></v-text-field>
                                <v-btn class="grey lighten-3 ma-3" @click="addTag">Add Tag</v-btn>
                            </v-row>
                            <v-textarea outlined label ="All tags" v-model="form.tags" rows="2" no-resize readonly></v-textarea>
                            <v-btn color="grey lighten-3" allign-right :disabled="!form.isFormValid" @click="submit">Publish Post</v-btn>
                        </v-form>
                    </v-col>
                    <v-col>
                        <img :src="imageUrl" height="500" width="400"/>
                    </v-col>
                    </v-row>
                    </v-container>
                </v-container>

    <v-footer app height="45px" class="grey lighten-3 justify-center">
        <v-container>
          <v-row justify="center">
            <v-btn class= "mx-2" @click="$router.push('/').catch(()=>{})">
              <v-icon>mdi-home</v-icon>
            </v-btn>

            <v-btn class= "mx-2" @click="$router.push('/search').catch(()=>{})">
              <v-icon>mdi-magnify</v-icon>
            </v-btn>

            <v-btn
              class= "mx-2"
              @click="$router.go()"
            >
              <v-icon>mdi-plus-box</v-icon>
            </v-btn>

            <v-btn class= "mx-2">
              <v-icon>mdi-bell-ring</v-icon>
            </v-btn>

            <v-btn
              class= "mx-2"
              @click="$router.push('/account').catch(()=>{})"
            >
              <v-icon>mdi-account</v-icon>
            </v-btn>
          </v-row>
        </v-container>
    </v-footer>
    </v-app>
</template>

<script>

import firebase from 'firebase/app'
import 'firebase/firebase-storage'
import { getUsername } from '../security/token.js'
import axios from 'axios'

export default {
    name: 'AddPostComment',
    components: {

    },
    data(){
        return {
            imageUrl: "",
            image: null,
            uploadValue: 0,
            firebaseURL: [],
            tag : "",
            form:{
                name: "",
                tags: [],
                location: "",
                isFormValid: false,
            },
            rules:{
                descriptionRules: [
                    description => !!description || 'Description is required!',
                    description => (description && description.length <= 100) || 'Description must contain less than 100 characters!',
                ],
                locationRules: [
                    description => !!description || 'Description is required!',
                    description => (description && description.length <= 30) || 'Description must contain less than 100 characters!',
                ],
            },
            snackbar: false,
            snackbarText: ""
        }
    },
    mounted() {

    },
    methods: {
        uploadFile(){
            this.$refs.fileInput.click();
        },
        onFilePicked(event){    
            const files = event.target.files;
            let filename = files[0].name;
            console.log(filename)
            if(filename.lastIndexOf('.') <= 0){
               return alert("Please upload valid image file!")
            }

            const fileReader = new FileReader();
            fileReader.addEventListener('load', () =>{
                this.imageUrl = fileReader.result;
            })
            fileReader.readAsDataURL(files[0])
            this.image = files[0]

            const storageRef = firebase.storage().ref(`${filename}`).put(this.image)
            storageRef.on(`state_changed`,snapshot=>{
                this.uploadValue = (snapshot.bytesTransferred/snapshot.totalBytes)*100;
                }, error=>{console.log(error.message)},
                ()=>{this.uploadValue=100;
                storageRef.snapshot.ref.getDownloadURL().then((url)=>{
                this.firebaseURL.push(url)
                console.log(this.firebaseURL)
                });
            });
        },
        addTag(){
            if(this.tag != "" && !this.form.tags.includes(this.tag))
                this.form.tags.push(this.tag);
            this.tag = "";
        },
        submit() {
            if(this.firebaseURL == ""){
                alert("Image is still not uploaded!")
            }else {
                let postUploadDTO = {
                description : this.form.description,
                tags : this.form.tags,
                mediaPaths : this.firebaseURL,
                uploadDate: new Date(),
                location: this.form.location,
                username: getUsername(),
                }
                axios.post('http://localhost:8081/api/media-content/new-post',
                    postUploadDTO
                ).then(response => {
                    console.log(response)
                    this.$router.push('/account').catch(()=>{})
                }).catch(error => {
                    if(error.response.status === 500){
                        this.snackbarText = "Internal server error occurred!";
                        this.snackbar = true;
                    }else if(error.response.status === 400){
                        this.snackbarText = "Bad request, try again!";
                        this.snackbar = true;
                    }
                })
            }
        }
    }
}
</script>
