<template>  
  <el-container class="home-container">  
    <div class="content-wrapper">  
      <h1 class="title">欢迎使用蓝鼠兑换</h1>  
      <p class="description">请选择上方的功能进行操作。</p>  
    </div>  
  </el-container>  
</template>  
  
<script setup lang="ts">   
</script>  
  
<style scoped>  
.home-container {  
  display: flex;  
  justify-content: center;  
  align-items: center;  
  height: 100vh; 
  background-color: #f5f5f5;
  padding: 20px; 
  box-sizing: border-box;
}  
  
.content-wrapper {  
  text-align: center;  
  max-width: 800px; 
}  
  
.title {  
  color: #333;  
  font-size: 36px;  
  font-weight: bold;  
  margin-bottom: 20px;   
}  
  
.description {  
  color: #666;  
  font-size: 18px; 
  line-height: 1.5;  
}  
</style>