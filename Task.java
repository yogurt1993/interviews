public abstract class Task {
   private String name; //
   private TaskStatus status;
   private long schedule;
   private Set<Task> dependsOn;

   public abstract void runTask();

   protected void finish(){
       status = TaskStatus.Finished;
       TaskManager.getInstance().taskFinished(this);
   }

   protected void reschedule(long timestamp){
       schedule = timestamp;
       TaskManager.getInstance().taskFinished(this);
       TaskManager.getInstance().skeduleTask(this); 
   }

   @Override
   public boolean equals(Object o) {
       if (this == o) return true;
       if (!(o instanceof Task)) return false;

       Task task = (Task) o;

       return schedule == task.schedule && name.equals(task.name);
   }

   @Override
   public int hashCode() {
       int result = name.hashCode();
       result = 31 * result + status.hashCode();
       result = 31 * result + (int) (schedule ^ (schedule >>> 32));
       return result;
   }

  // (get/set all fields ommited)
}
