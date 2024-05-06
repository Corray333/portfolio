<script setup>

const props = defineProps(['project'])

</script>

<template>
    <article class=" portfolio_card pointer grid grid-cols-2 gap-5 bg-white z-50 relative p-5 max-w-[900px] drop-shadow-md hover:drop-shadow-2xl duration-500">
        <img :src="`/images/home/${project.img}`" alt="" class=" w-full">
        <div class="info w-full flex flex-col gap-2">
            <h3 class="text-xl">{{ project.title }}</h3>
            <p class="text-sm">{{ project.description }}</p>
            <div class="flex gap-2">
                <Icon v-for="(icon, i) of project.stack" :key="i" :name="icon" class="text-[3rem]"/>
            </div>
        </div>
    </article>
</template>


<style scoped>

.portfolio_card:hover{
    transform: translateZ(250px);
    --webkit-font-smoothing: antialiased;
}

</style>