<template>
 
   <card-component
     title="Login"
     icon="lock"

   >
     <form
   
       method="POST"
       @submit.prevent="submit"
     >
       <b-field label="E-mail Address">
         <b-input
           v-model="form.email"
           name="email"
           type="email"
           required
         />
       </b-field>
 
       <b-field label="Password">
         <b-input
           v-model="form.password"
           type="password"
           name="password"
           required
         />
       </b-field>
 
       <b-field>
         <b-checkbox
           v-model="form.remember"
           type="is-black"
           class="is-thin"
         >
           Remember me
         </b-checkbox>
       </b-field>
 
       <hr>
 
       <b-field grouped>
         <div class="control">
           <b-button
             native-type="submit"
             type="is-black"
             :loading="isLoading"
           >
             Login
           </b-button>
         </div>
       </b-field>
     </form>
   </card-component>

 </template>
 
 <script>
 import CardComponent from '@/components/CardComponent.vue'
 
 export default {
   name: 'Login',
   components: { CardComponent },
   data () {
     return {
       isLoading: false,
       form: {
         email: 'john.doe@example.com',
         password: 'my-secret-password-9e9w',
         remember: false
       }
     }
   },
   methods: {
     submit () {
       this.isLoading = true
 
       setTimeout(() => {
         this.isLoading = false
          this.$router.push('/')
       }, 750)
     }
   }
 }
 </script>
 <style>
#form{
  display: flex;
  align-items: center;
  justify-content: center;
}
</style>