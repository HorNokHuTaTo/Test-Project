<template>
  <div class="exam-container">
    <div class="header">
      {{ result ? "IT 10-2" : "IT 10-1" }}
    </div>
    <div class="form-row">
      <label>ชื่อ-สกุล</label>
      <input v-model="name" placeholder="กรอกชื่อผู้สอบ" />
    </div>

    <div v-for="(q, qi) in questions" :key="qi" class="question-box">
      <div class="question-text">{{ q.question_text }}</div>

      <div v-for="(c, ci) in q.choices" :key="ci" class="choice">
        <label>
          <input
            type="radio"
            :name="'q' + qi"
            :value="ci"
            v-model="answers[qi]"
            :disabled="mode === 'result'"
          />
          {{ c }}
        </label>
      </div>
    </div>

    <button class="btn-submit" @click="handleSubmit">
  {{ submitted ? 'สอบอีกครั้ง' : 'ส่งข้อสอบ' }}
</button>
    <p v-if="errorMessage" style="color:red">{{ errorMessage }}</p>
    <div v-if="mode === 'result'" class="result-box">
  คุณ {{ result.name }} สอบได้คะแนน : {{ result.score }} / {{ result.total }}
</div>
<div v-if="submitted" class="result-text">
  คุณ {{ result.name }} สอบได้คะแนน : {{ result.score }}/{{ result.total }}
</div>
  </div>
</template>

<script>
const API = "https://test-project-0w71.onrender.com/api/it10";

export default {
  name: "It10",
  data() {
    return {
      name: "",
      questions: [],
      answers: [],
      result: null,
      submitted: false,
      errorMessage: "",
    };
  },

  mounted() {
    this.loadQuestions();
  },

  methods: {
    async loadQuestions() {
      const res = await fetch(API + "/questions");
      this.questions = await res.json();
      this.answers = new Array(this.questions.length).fill(null);
    },

    async handleSubmit() {
      if (!this.submitted) {
        const payload = {
          name: this.name,
          answers: this.answers,
        };

        const res = await fetch(API + "/submit-exam", {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify(payload),
        });

        if (!res.ok) {
          const err = await res.json();
          alert(err.error || "Submit failed");
          return;
        }

        this.result = await res.json();
        this.submitted = true;
      } else {
        this.name = "";
        this.answers = [];
        this.result = null;
        this.submitted = false;
        await this.loadQuestions();
      }
    },
  },
};
</script>


<style scoped>
.exam-container {
  width: 800px;
  margin: auto;
  border: 1px solid #333;
  padding: 20px;
  position: relative;
}

.header {
  background: #3cb054;
  text-align: center;
  padding: 10px;
  font-weight: bold;
}

.form-row {
  margin: 15px 0;
}

.question-box {
  margin: 10px 0;
  padding: 10px;
}

.choice {
  margin-left: 15px;
}

.submit-btn {
  margin-top: 15px;
}

.result-bottom {
  position: absolute;
  right: 20px;
  bottom: 20px;
  font-weight: bold;
}
.result-text {
  margin-top: 20px;
  text-align: right;      
  font-weight: bold;     
  font-size: 16px;
}

.btn-submit {
  background: #2e6fd3;
  color: #fff;
  border: none;
  padding: 8px 16px;
  cursor: pointer;
}

</style>
