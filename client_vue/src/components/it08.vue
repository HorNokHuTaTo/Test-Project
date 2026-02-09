<template>
  <div class="page">
    <div v-if="page === 'list'">
      <div class="header">IT 08-1</div>

      <button class="btn-add" @click="goAdd">เพิ่มข้อสอบ</button>

      <div
        v-for="(q, index) in questions"
        :key="q.id"
        class="question-block"
      >
        <div class="q-title">
          {{ index + 1 }}. {{ q.question_text }}
          <button class="btn-del" @click="deleteQuestion(q.id)">ลบ</button>
        </div>

        <div class="choices">
          <label v-for="(c, i) in q.choices" :key="i">
        <input type="radio" disabled />
        {{ c }}
        </label>

        </div>
      </div>
    </div>
    <div v-else>
      <div class="header">IT 08-2</div>

      <div class="form">
        <label>คำถาม</label>
        <input v-model="form.question_text" />

        <div v-for="(c, i) in form.choices" :key="i">
  <input v-model="form.choices[i]" :placeholder="`คำตอบ ${i + 1}`" />
</div>
        <div class="actions">
          <button class="btn-save" @click="saveQuestion">บันทึก</button>
          <button class="btn-cancel" @click="page = 'list'">ยกเลิก</button>
        </div>
      </div>
    </div>

  </div>
</template>


<script setup>
import { ref, onMounted } from "vue"

const API = "https://test-project-0w71.onrender.com/api/it08"

const page = ref("list")
const questions = ref([])

const form = ref({
  question_text: "",
  choices: ["", "", "", ""],
})

onMounted(loadQuestions)

async function loadQuestions() {
  const res = await fetch(`${API}/questions`)
  questions.value = await res.json()
}

function goAdd() {
  page.value = "add"
}

async function saveQuestion() {
  const res = await fetch(`${API}/insert-question`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(form.value),
  })

  if (!res.ok) {
    alert("บันทึกไม่สำเร็จ")
    return
  }

  resetForm()
  page.value = "list"
  await loadQuestions()
}

async function deleteQuestion(id) {
  if (!confirm("ต้องการลบข้อสอบนี้หรือไม่ ?")) return

  const res = await fetch(`${API}/delete-question?id=${id}`, {
    method: "DELETE",
  })

  if (!res.ok) {
    alert("ลบไม่สำเร็จ")
    return
  }

  await loadQuestions()
}

function resetForm() {
  form.value = {
    question_text: "",
    choices: ["", "", "", ""],
  }
}
</script>

<style scoped>
.page {
  width: 900px;
  margin: 20px auto;
  border: 1px solid #333;
}
.header {
  background: #00b050;
  text-align: center;
  padding: 8px;
  font-weight: bold;
}
.btn-add {
  background: #6aa84f;
  color: #fff;
  border: none;
  padding: 6px 14px;
  margin: 10px;
  cursor: pointer;
}
.question-block {
  padding: 10px 20px;
}
.q-title {
  font-weight: bold;
}
.btn-del {
  background: red;
  color: white;
  border: none;
  padding: 2px 8px;
  margin-left: 10px;
  cursor: pointer;
}
.choices label {
  display: block;
  margin-left: 20px;
}
.form {
  padding: 20px;
}
.form input {
  width: 100%;
  margin-bottom: 10px;
  padding: 6px;
}
.actions {
  text-align: center;
}
.btn-save {
  background: #2e6fd3;
  color: #fff;
  padding: 6px 16px;
  border: none;
    cursor: pointer;
}
.btn-cancel {
  background: red;
  color: #fff;
  padding: 6px 16px;
  border: none;
  margin-left: 10px;
    cursor: pointer;
}
</style>
