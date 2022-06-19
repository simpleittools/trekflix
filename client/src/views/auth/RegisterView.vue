<template>
  <div class="form-container">
    <div class="head-container">
<!--      replace with database logo file-->
      <img class="logo-image" src="../../assets/img/TrekFlix.png" alt="">
      <h2 class="page-title">Sign Up</h2>
      <p class="page-subtitle">
        to continue to TrekFlix
      </p>
    </div>

    <div class="mt-8 sm:mx-auto sm:w-full sm:max-w-md">
      <div class="bg-white py-8 px-4 shadow sm:rounded-lg sm:px-10">
        <VeeForm class="space-y-6" @submit="handleSubmit">
          <FirstNameComponent v-model="data.firstName" />
          <LastNameComponent v-model="data.lastName" />
          <UsernameComponent v-model="data.username" />
          <EmailFormComponent v-model="data.email" />
          <PasswordFormComponent v-model="data.password" />
          <PasswordValidateFormComponent v-model="data.password2" />

          <div class="button-container">
            <button type="submit" class="button-submit">Submit</button>
            <button type="reset" class="button-reset" @click.prevent="handleClear">Clear</button>
          </div>
        </VeeForm>
      </div>
    </div>
  </div>

</template>

<script>
import {reactive} from "vue";
import {Form, Field, ErrorMessage} from 'vee-validate'
import BaseInput from "@/components/formComponents/BaseInput";
import FirstNameComponent from "@/components/formComponents/FirstNameComponent";
import LastNameComponent from "@/components/formComponents/LastNameComponent";
import UsernameComponent from "@/components/formComponents/UsernameComponent";
import EmailFormComponent from "@/components/formComponents/EmailFormComponent";
import PasswordFormComponent from "@/components/formComponents/PasswordFormComponent";
import PasswordValidateFormComponent from "@/components/formComponents/PasswordValidateFormComponent";
export default {
  name: "RegisterView",
  components: {
    VeeForm: Form,
    VeeField: Field,
    VeeErrorMessage: ErrorMessage,
    BaseInput,
    FirstNameComponent,
    LastNameComponent,
    UsernameComponent,
    EmailFormComponent,
    PasswordFormComponent,
    PasswordValidateFormComponent,
  },
  props:['firstName'],
  setup() {

    const data = reactive({
      firstName: '',
      lastName: '',
      username: '',
      email: '',
      email2: '',
      password: '',
      password2: '',
    })


    const handleSubmit = () => {
      data.firstName
      data.lastName
      data.username
      data.email
      data.email2
      data.password
      data.password2
      console.log(data)
    }
    const handleClear = () => {
      data.firstName = ''
      data.lastName = ''
      data.username = ''
      data.email = ''
      data.email2 = ''
      data.password = ''

      data.password2 = ''
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