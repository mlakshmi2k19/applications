<?php

@include 'config.php';

session_start();

$user_id = $_SESSION['user_id'];

if(!isset($user_id)){
   header('location:login.php');
}

?>

<!DOCTYPE html>
<html lang="en">
<head>
   <meta charset="UTF-8">
   <meta http-equiv="X-UA-Compatible" content="IE=edge">
   <meta name="viewport" content="width=device-width, initial-scale=1.0">
   <title>about</title>

   <!-- font awesome cdn link  -->
   <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.1.1/css/all.min.css">

   <!-- custom css file link  -->
   <link rel="stylesheet" href="css/style3.css">
   <link rel="stylesheet" href="css/sidenav.css">

</head>
<body>
   
<?php include 'header.php'; ?>

<section class="about">

   <div class="row">

      <div class="box">
         <img src="images/about-img-1.png" alt="">
         <h3>why choose us?</h3>
         <p>Research finds that people who eat home-cooked meals on a regular basis tend to be happier and healthier and consume less sugar and processed foods, which can result in higher energy levels and better mental health. Eating home-cooked meals five or more days a week is even associated with a longer life.</p>
         <a href="contact.php" class="btn">contact us</a>
      </div>

      <div class="box">
         <img src="images/about-img-2.png" alt="">
         <h3>what we provide?</h3>
         <p>We are located in the city center. The highest standards of service.We provides platform for anyone who knows cooks,who want to cook,especially for women's who staying home and also want to make them financially independant.  </p>
         <a href="shop.php" class="btn">our shop</a>
      </div>

   </div>

</section>

<section class="reviews">

   <h1 class="title">customer reivews</h1>

   <div class="box-container">

      <div class="box">
         <img src="images/pic-1.png" alt="">
         <p>Food is good and reliable, fresh and really homemade.
I've tried lunch and breakfast, always more than ok.Thanks to Saravanas.</p>
         <div class="stars">
            <i class="fas fa-star"></i>
            <i class="fas fa-star"></i>
            <i class="fas fa-star"></i>
            <i class="fas fa-star"></i>
            <i class="fas fa-star-half-alt"></i>
         </div>
         <h3>johnson</h3>
      </div>

      <div class="box">
         <img src="images/pic-2.png" alt="">
         <p>I like specially murugas chai... as simple as that.
Nice, quiet place, try it if you have the chance. Affordable and really good value for money..</p>
         <div class="stars">
            <i class="fas fa-star"></i>
            <i class="fas fa-star"></i>
            <i class="fas fa-star"></i>
            <i class="fas fa-star"></i>
            <i class="fas fa-star-half-alt"></i>
         </div>
         <h3>joheo</h3>
      </div>

      <div class="box">
         <img src="images/pic-3.png" alt="">
         <p>Edmond can cook dal fry as well as french fries and fish, which is nice for 
those visiting Goa that dislike or get soon tired of the spicy Indian tastes..</p>
         <div class="stars">
            <i class="fas fa-star"></i>
            <i class="fas fa-star"></i>
            <i class="fas fa-star"></i>
            <i class="fas fa-star"></i>
            <i class="fas fa-star-half-alt"></i>
         </div>
         <h3>Alice</h3>
      </div>

      <div class="box">
         <img src="images/pic-4.png" alt="">
         <p> Food is good and reliable, fresh and really homemade.
I've tried lunch and breakfast, always more than ok.Thanks to Saravanas.</p>
         <div class="stars">
            <i class="fas fa-star"></i>
            <i class="fas fa-star"></i>
            <i class="fas fa-star"></i>
            <i class="fas fa-star"></i>
            <i class="fas fa-star-half-alt"></i>
         </div>
         <h3>Bob</h3>
      </div>
      <div class="box">
         <img src="images/pic-6.png" alt="">
         <p>Edmond can cook dal fry as well as french fries and fish, which is nice for 
those visiting Goa that dislike or get soon tired of the spicy Indian tastes..</p>
         <div class="stars">
            <i class="fas fa-star"></i>
            <i class="fas fa-star"></i>
            <i class="fas fa-star"></i>
            <i class="fas fa-star"></i>
            <i class="fas fa-star-half-alt"></i>
         </div>
         <h3>Michael</h3>
      </div>

      <div class="box">
         <img src="images/pic-5.png" alt="">
         <p>I like specially murugas chai... as simple as that.
Nice, quiet place, try it if you have the chance. Affordable and really good value for money.</p>
         <div class="stars">
            <i class="fas fa-star"></i>
            <i class="fas fa-star"></i>
            <i class="fas fa-star"></i>
            <i class="fas fa-star"></i>
            <i class="fas fa-star-half-alt"></i>
         </div>
         <h3>john deo</h3>
      </div>

      

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