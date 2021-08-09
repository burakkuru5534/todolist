<template>
  <div id="app" class="container-fluid text-center">
    <h1 class="text-info">{{ title }}</h1>
    <p>{{ msg }}</p>

    <div class="col">
      <div v-if="error" class="alert alert-danger" @click="error = !error">
        <strong>Error:</strong> Please add the task name first!
      </div>
      <form @submit.prevent="addTask">
        <div class="input-group mb-3">
          <input type="text" class="form-control" placeholder="Task Name" v-model="EventContent">
          <div class="input-group-append">
            <button v-if="update" class="btn btn-warning" type="submit">
              <i class="fa fa-pencil" aria-hidden="true"></i>
            </button>
            <button v-else class="btn btn-success" type="submit">
              <i class="fa fa-plus" aria-hidden="true"></i>
            </button>
          </div>
        </div>
      </form>

      <button v-if="deleteMultiple" class="btn btn-danger mb-3" @click="deleteMulti">Delete Selected</button>

      <ul class="list-group">
        <li v-for="(task_name,index) in tasks" :key="index" class="list-group-item list-group-item-info">
          <div class="row">
            <div class="col" @click="loadData(index)">{{ task_name['task'] }}</div>
          </div>
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
export default {
  name: 'app',
  data () {

    return {
      title: 'To do List App',
      msg: 'Add your to do \'s',
      EventContent: "",
      tasks: [

      ],
      error: false,
      update: false,
      updateIndex: null,
      ids: [],
      deleteMultiple: false
    }
  },
  // This is run whenever the page is loaded to make sure we have a current task list
  created: function() {
    // Use the vue-resource $http client to fetch data from the /tasks route
    this.$http.get('/todo').then(function(response) {
      this.tasks = response.data.items ? response.data.items : []
    })
  },

  methods: {
    createTask: function() {
      if (!$.trim(this.newTask.name)) {
        this.newTask = {}
        return
      }

      // Post the new task to the /tasks route using the $http client
      this.$http.put('/tasks', this.newTask).success(function(response) {
        this.newTask.id = response.created
        this.tasks.push(this.newTask)
        console.log("Task created!")
        console.log(this.newTask)
        this.newTask = {}
      }).error(function(error) {
        console.log(error)
      });
    },

    deleteTask: function(index) {
      // Use the $http client to delete a task by its id
      this.$http.delete('/tasks/' + this.tasks[index].id).success(function(response) {
        this.tasks.splice(index, 1)
        console.log("Task deleted!")
      }).error(function(error) {
        console.log(error)
      })
    }
  }
}
</script>

<style>
</style>
