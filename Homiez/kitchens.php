<?php

@include 'config.php';

session_start();

$user_id = $_SESSION['user_id'];

if(!isset($user_id)){
   header('location:login.php');
};

?>



<!DOCTYPE html>
<html lang="en">
<head>
   <meta charset="UTF-8">
   <meta http-equiv="X-UA-Compatible" content="IE=edge">
   <meta name="viewport" content="width=device-width, initial-scale=1.0">
   <title>Kitchen</title>

   <!-- font awesome cdn link  -->
   <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.1.1/css/all.min.css">

   <!-- custom css file link  -->
   <link rel="stylesheet" href="css/style4.css">
   <link rel="stylesheet" href="css/sidenav.css">

</head>
<body>
   
<?php include 'header.php'; ?>
<section class="products">

   <h1 class="title">All kitchens</h1>

   <div class="box-container">

   <?php
      $select_kitchens = $conn->prepare("SELECT * FROM `kitchens`");
      $select_kitchens->execute();
      if($select_kitchens->rowCount() > 0){
         while($fetch_kitchens = $select_kitchens->fetch(PDO::FETCH_ASSOC)){ 
            $select_user = $conn->prepare("SELECT * FROM `users` where id=? LIMIT 1");
            $select_user->execute([$fetch_kitchens['user_id']]);
            $fetch_user = $select_user->fetch(PDO::FETCH_ASSOC)
            // while($fetch_user = $select_user->fetch(PDO::FETCH_ASSOC)){ 
   ?>
   <form action="" class="box" method="POST">
      
      <div class="name" style="font-weight:bold"><?= $fetch_kitchens['kname']; ?></div>
      <br><br>
      <img src="uploaded_img/<?= $fetch_user['image']; ?>" alt=""><br><br><br>
      <a href="kit_prod.php?kname=<?= $fetch_kitchens['kname']; ?>" class="btn">visit</a>
   </form>
   <?php
      }
   }else{
      echo '<p class="empty">No kitchens added yet!</p>';
   }
   ?>

   </div>

</section>





<?php include 'footer.php'; ?>

<script src="js/script.js"></script>
<script>
    const menuToggle=document.querySelector('.menuToggle');
    const navigation=document.querySelector('.navigation');
    menuToggle.onclick=function(){
        navigation.classList.toggle('open');
    }
</script>

</body>
</html>