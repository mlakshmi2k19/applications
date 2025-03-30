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
   <title>home page</title>

   <!-- font awesome cdn link  -->
   <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.1.1/css/all.min.css">
   

   <!-- custom css file link  -->
   <link rel="stylesheet" href="css/style4.css">
   <link rel="stylesheet" href="css/sidenav.css">
   
</head>
<body>
   
<?php include 'header.php'; ?>
<section class="home-category">

   <h1 class="title">In the spotlight!</h1>

   <div class="box-container" style="gap:4.5rem;">
   <?php 
      $select_cat = $conn->prepare("SELECT * FROM `category_list`");
      $select_cat->execute();
      if($select_cat->rowCount() > 0){
         while($fetch_cat = $select_cat->fetch(PDO::FETCH_ASSOC)){ 
   ?>
      <div class="box">
         <img src="uploaded_img/<?= $fetch_cat['image']; ?>" alt=""><br><br>
         <a href="category.php?category=<?= $fetch_cat['id']; ?>" class="btn"><?= $fetch_cat['name'];?></a>
      </div>
      <?php 
            }
         }else{
            echo '<p class="empty">No categories added yet!</p>';
         }
      ?>

   </div>

</section>
<br><br>
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