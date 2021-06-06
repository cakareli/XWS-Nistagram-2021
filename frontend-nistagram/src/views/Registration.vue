<template>
<nav></nav>
    <!-- <div id="app">
        <div class="app-phone">
            <div class="logo">              
                <img src="../assets/logoedited2.png">
            </div> 
            <div class="registration-backgorund">
                <form id="registration-input" @submit="checkForm">
                    <div class="row">
                        <div class="column">
                            <img src="../assets/profile2.png" height="35" width="35">
                            <br><br><br><br><br><br>
                            <img src="../assets/password.png" height="35" width="35">
                            <br><br><br>
                            <img src="../assets/email.png" height="25" width="35">
                            <br><br><br>
                            <img src="../assets/phone.png" height="30" width="30">
                            <br><br>
                            <img src="../assets/gender.png" height="50" width="50">
                            <br><br><br>
                            
                        </div>
                        <div class="column">
                            <input placeholder="Name" type="text">
                            <br><br>
                            <input placeholder="Surname" type="text">
                            <br><br>
                            <input placeholder="Username" type="text">
                            <br><br>
                            <input placeholder="Password" type="password">
                            <br><br>
                            <input placeholder="Repeat password" type="password">
                            <br><br>
                            <input placeholder="Email" type="email">
                            <br><br><br>
                            <input placeholder="Phone Number" type="number">
                            <br><br><br><br>
                            <input type="radio" id="male" name="gender" value="male" checked>
                            <label for="male">Male</label><br>
                            <input type="radio" id="female" name="gender" value="female">
                            <label for="female">Female</label><br>
                            <br><br><br>                         
                        </div>
                        <div class="column">
                            <img src="../assets/birthday.png" height="35" width="35">
                            <br><br><br>
                            <label>User type:</label>
                            <br><br><br><br>
                            <img src="../assets/website.png" height="35" width="35">
                            <br><br>
                            <img src="../assets/bio.png" height="45" width="45">
                        </div>
                        <div class="column">
                            <input placeholder="Birthday date" type="date" height="35">
                            <br><br><br><br>
                            <select>
                                <option disabled value="">Please select one</option>
                                <option>Regular User</option>
                                <option>Agent</option>
                            </select> 
                            <br><br><br><br><br>
                            <input type="url" placeholder="URL"> 
                            <br><br><br>
                            <textarea name="bio" id="bio" cols="30" rows="10" placeholder="Biography"></textarea>
                        </div>
                    </div>
                    <div>
                        <button class="submit">Submit</button>
                    </div>
                </form>
            </div>
            <div class="lower-toolbar">
                <button @click="$router.push('/')">
                    <img src="../assets/back.png">
                </button>
            </div>
        </div>
    </div> -->
</template>

<script>
export default {
    name: 'Registration',
    data(){
        return {
            form: {

            },
            rules: {

            }
        }
    }
}
</script>


<style lang="scss" src="../style/home.scss">

</style>