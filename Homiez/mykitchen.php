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
   <title>category</title>

   <!-- font awesome cdn link  -->
   <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.1.1/css/all.min.css">

   <!-- custom css file link  -->
   <link rel="stylesheet" href="css/style4.css">
   <link rel="stylesheet" href="css/sidenav.css">

</head>
<body>
   
<?php include 'header.php'; ?>

<section class="products">

   <h1 class="title">My kitchen</h1>

   <div class="box-container">

   <?php
      $select_kitchen = $conn->prepare("SELECT * FROM `kitchens` WHERE user_id = ?");
      $select_kitchen->execute([$user_id]);
      if($select_kitchen->rowCount() > 0){
         while($fetch_kitchen = $select_kitchen->fetch(PDO::FETCH_ASSOC)){ 
            $select_products = $conn->prepare("SELECT * FROM `products` WHERE kname = ?");
            $select_products->execute([$fetch_kitchen['kname']]);
            if($select_products->rowCount() > 0){
                while($fetch_products = $select_products->fetch(PDO::FETCH_ASSOC)){ 
   ?>
   <form action="" class="box" method="POST">
      <div class="price">$<span><?= $fetch_products['price']; ?></span>/-</div>
      <a href="view_page.php?pid=<?= $fetch_products['id']; ?>" class="fas fa-eye"></a><br><br><br>
      <img src="uploaded_img/<?= $fetch_products['image']; ?>" alt="">
      <div class="name"><?= $fetch_products['pname']; ?></div>
      </form>
   <?php
         }
      }else{
         echo '<p class="empty">No products available!</p>';
      }
    }
    }else{
        echo '<p class="empty">No kitchen available to this user id!</p>';
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