<script setup>
import {ref} from 'vue'
import axios from 'axios'
import {useRouter} from 'vue-router'

const router = useRouter()

const login = ref("")
const password = ref("")

const submit = async ()=>{
    try{
        let {data} = await axios.post('/api/admin/login',{
            login:login.value,
            password:password.value
        })
        console.log(data)
        localStorage.setItem('Authorization',data.authorization)
        localStorage.setItem('Refresh',data.refresh)
        router.push(`/admin`)
    } catch(err){
        console.log(err)
    }
}

</script>

<template>
    <section class="w-full h-screen flex flex-col justify-center items-center">
        <h1 class="text-2xl">Login</h1>
        <div class="flex flex-col gap-2">
            <div class="flex">
                <label for="">Login: </label>
                <input v-model="login" type="text" class=" border-b-2 border-dark z-50 pointer w-full">
            </div>
            <div class="flex">
                <label for="">Password: </label>
                <input v-model="password" type="password" class=" border-b-2 border-dark z-50 pointer w-full">
            </div>
            <button @click="submit" type="button" class="mx-auto w-fit flex border-2 border-dark bg-white relative z-50 pointer btn"><a class="bg-white px-4 py-2">submit</a></button>
        </div>
    </section>
</template>


<style scoped>

.btn>a::after{
    padding: 0.5rem 1rem;
}

</style>