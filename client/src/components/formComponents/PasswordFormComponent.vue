<template>
  <div class="mt-1 sm:mt-0 sm:col-span-2">
    <div class="input-frame">
      <span class="input-label"> Password </span>
      <VeeField type="password" :value="modelValue"
             @input="$emit('update:modelValue', $event.target.value)"
             name="password" :rules="(value) => {
               if(value && value.trim()) return true
               return 'A password is required'
             }"

             autocomplete="password" class="input-field"
      />
    </div>
    <VeeErrorMessage name="password" class="form-error"/>
  </div>
</template>

<script>
import {ErrorMessage, Field} from "vee-validate";

export default{
  components: {
    VeeField: Field,
    VeeErrorMessage: ErrorMessage,

  },
  }
</script>

<script setup>
  defineProps({
    label: {
      type: String,
      default: "",
    },
    modelValue: {
      type: String,
      default: "",
    }
  })
</script>

<style scoped>

</style>