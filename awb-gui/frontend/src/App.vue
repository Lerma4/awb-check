<script setup lang="ts">
import { ref } from 'vue'
// @ts-ignore
import { CalculateCIN, GenerateRandomAWB } from '../wailsjs/go/main/App'

const awbInput = ref('')
const result = ref('')
const resultValue = ref('')
const error = ref('')

const calculate = async () => {
  error.value = ''
  result.value = ''
  try {
    const cin = await CalculateCIN(awbInput.value)
    result.value = `CIN: ${cin} -> Full AWB: ${awbInput.value}${cin}`
    resultValue.value = `${awbInput.value}${cin}`
  } catch (e: any) {
    error.value = e as string
  }
}

const generate = async () => {
  error.value = ''
  try {
    const awb = await GenerateRandomAWB()
    result.value = `Generated AWB: ${awb}`
    awbInput.value = awb.substring(0, 7)
    resultValue.value = awb
  } catch (e: any) {
    error.value = e as string
  }
}
const copied = ref(false)

const copyToClipboard = async () => {
  if (resultValue.value) {
    try {
      await navigator.clipboard.writeText(resultValue.value)
      copied.value = true
      setTimeout(() => copied.value = false, 2000)
    } catch (err) {
      console.error('Failed to copy:', err)
    }
  }
}
</script>

<template>
  <div class="min-h-screen bg-gray-100 flex items-center justify-center p-4 font-sans">
    <div class="bg-white p-8 rounded-xl shadow-2xl w-full max-w-md">
      <h1 class="text-3xl font-bold text-gray-800 mb-6 text-center">AWB Checker</h1>
      
      <div class="space-y-4">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">AWB Number (7 digits)</label>
          <input 
            v-model="awbInput" 
            type="text" 
            maxlength="7"
            class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 outline-none transition-all"
            placeholder="1234567"
            @keyup.enter="calculate"
          />
        </div>

        <div class="flex gap-3">
          <button 
            @click="calculate" 
            class="flex-1 bg-blue-600 hover:bg-blue-700 text-white font-semibold py-2 px-4 rounded-lg transition-colors duration-200"
          >
            Calculate CIN
          </button>
          <button 
            @click="generate" 
            class="flex-1 bg-gray-800 hover:bg-gray-900 text-white font-semibold py-2 px-4 rounded-lg transition-colors duration-200"
          >
            Random AWB
          </button>
        </div>

        <div v-if="result" class="mt-6 p-4 bg-green-50 border border-green-200 rounded-lg text-green-800 font-medium flex justify-between items-center">
          <span>{{ result }}</span>
          <button 
            @click="copyToClipboard" 
            class="ml-4 px-3 py-1 bg-green-600 hover:bg-green-700 text-white text-sm rounded transition-colors duration-200"
          >
            {{ copied ? 'Copied!' : 'Copy' }}
          </button>
        </div>

        <div v-if="error" class="mt-6 p-4 bg-red-50 border border-red-200 rounded-lg text-red-800 text-center font-medium">
          {{ error }}
        </div>
      </div>
    </div>
  </div>
</template>
