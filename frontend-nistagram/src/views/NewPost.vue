<template>
    <v-container>
        <v-row>
            <v-form>
                <v-btn raised class="grey lighten-3" @click="uploadFile">Upload Image</v-btn>
                <input type="file" style="display: none" ref="fileInput" accept="image/*" @change="onFilePicked"/>
            </v-form>
        </v-row>
        <v-row>
            <img :src="imageUrl" height="500" width="400"/>
        </v-row>
    </v-container>
</template>

<script>

import firebase from 'firebase/app'
import 'firebase/firebase-storage'

export default {
    name: 'AddPostComment',
    components: {

    },
    data(){
        return {
            imageUrl: "",
            image: null,
            uploadValue: 0,
            firebaseURL: "",
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
                this.firebaseURL =url;
                console.log(this.firebaseURL)
                });
            });
            
        }
    }
}
</script>
