<template>
  <div :class="[task.reminder ? 'reminder' : '', 'task']">
    <tabel>
      <tr>
        <th>{{ task.task }}</th>
        <th>{{ task.shop }}</th>
        <th>{{ task.product }}</th>
        <th>{{ task.size }}</th>
        <th>{{ task.color }}</th>
        <th>{{ task.status }}</th>
        <i @click="onDelete(task.id)" class="fas fa-times"></i>
      </tr>
    </tabel>
  </div>
</template>

<script>
export default {
  name: "Task-el",
  props: {
    task: Object,
  },
  methods: {
    onDelete(id) {
      this.$emit("delete-task", id);
    },
  },
};
</script>

<style scope>
.fas {
  color: red;
}
.task {
  background: rgb(55, 68, 92);
  height: 45px;
  width: 930px;
  margin-left: 70px;
  margin-right: 70px;
  margin-top: 5px;
  margin-bottom: 5px;
  padding: 10px 20px;
  cursor: pointer;
  border-radius: 10px;
}
.task.reminder {
  border-left: 5px solid green;
}

.task table {
  font-family: Arial, Helvetica, sans-serif;
  border-collapse: collapse;
}
.task th {
  text-align: left;
  min-width: 140px;
  max-width: 140px;
  height: 25px;
}
</style>
