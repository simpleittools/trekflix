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
          <UsernameComponent v-model="data.username"/>
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
import {Form} from 'vee-validate'
import UsernameComponent from "@/components/formComponents/UsernameComponent";
import PasswordFormComponent from "@/components/formComponents/PasswordFormComponent";

export default {
  name: "LoginView",
  components: {
    VeeForm: Form,
    UsernameComponent,
    PasswordFormComponent,
  },
  setup() {
    const data = reactive({
      username: '',
      password: '',
    })

    const handleSubmit = () => {
      data.username
      data.password

      console.log(data)
    }
    const handleClear = () => {
      data.username = ''
      data.password = ''
      console.log(data)
    }

    return {
      data,
      handleSubmit,
      handleClear
    }
  }
}
</script>

<style scoped>

</style>