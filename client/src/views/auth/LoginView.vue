<template>
  <div class="form-container">
    <div class="head-container">
      <!--      replace with database logo file-->
      <img class="logo-image" src="../../assets/img/TrekFlix.png" alt="">
      <h2 class="page-title">Login</h2>
    </div>

    <div class="mt-8 sm:mx-auto sm:w-full sm:max-w-md">
      <div class="bg-white py-8 px-4 shadow sm:rounded-lg sm:px-10">
        <VeeForm class="space-y-6" @submit="handleSubmit">
          <p v-if="loginError">Username or password failed.</p>
          <UsernameComponent v-model="data.username" :class="{'error-flag': loginError}"/>
          <PasswordFormComponent v-model="data.password" />

          <div class="button-container">
            <button type="submit" class="button-submit">Submit</button>
            <button type="reset" class="button-reset" @click="handleClear">Clear</button>
          </div>
        </VeeForm>
      </div>
    </div>
  </div>
</template>

<script>
import {reactive} from "vue";
import {useRouter} from "vue-router";
import {Form} from 'vee-validate'
import UsernameComponent from "@/components/formComponents/UsernameComponent";
import PasswordFormComponent from "@/components/formComponents/PasswordFormComponent";
import axios from "axios";
import {ref} from "vue";

export default {
  name: "LoginView",
  components: {
    VeeForm: Form,
    UsernameComponent,
    PasswordFormComponent
  }
}
</script>
<script setup>
    const router = useRouter()
    const loginError = ref(false)
    const data = reactive({
      username: '',
      password: '',
    })

    const handleSubmit = async () => {
      try {
        await axios.post('auth/login', {
          username: data.username,
          password: data.password,
        });
        await router.push({name: 'home'})
      } catch(err) {
        if (err.response) {
          loginError.value = true

        }
        console.log(err)
      }
    }
    const handleClear = () => {
      data.username = ''
      data.password = ''
      console.log(data)
    }


</script>

<style scoped>

</style>