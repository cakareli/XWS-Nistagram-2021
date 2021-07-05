<template>
    <v-dialog v-model="reportPostDialog" width = "450" height="500" persistent>
        <v-layout justify-center>
             <v-flex>
                <v-card color="grey lighten-4" class="pa-3">
                    <v-card-title>
                            <h3>Why are you reporting this post?</h3>
                    </v-card-title>
                        <v-radio-group>
                            <v-radio label="It's spam" value="It's spam"></v-radio>
                            <v-radio label="Nudity or sexual activity" value="Nudity or sexual activity"></v-radio>
                            <v-radio label="I just don't like it" value="I just don't like it"></v-radio>
                            <v-radio label="Hate speech or symbols" value="Hate speech or symbols"></v-radio>
                            <v-radio label="False information" value="False information"></v-radio>
                        </v-radio-group>
                    <v-card-actions>
                        <v-btn color="error" @click.native="close">Cancel</v-btn>
                        <v-spacer></v-spacer>
                        <v-btn color="success">Submit</v-btn>
                    </v-card-actions>
                </v-card>
             </v-flex>
        </v-layout>
    </v-dialog>
</template>

<script>

export default {
   name: 'ReportPost',
    components: {
    },
    props: {
        reportPostDialog: {
        default: false,
        },
    },
    data() {
        return {
            
        }
    },
    methods : {
        close() {
        this.$emit('update:reportPostDialog', false)
        },
    }
}
</script>