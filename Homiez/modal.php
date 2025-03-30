<?php
session_start();
include 'config.php';

$user_id = $_SESSION['user_id'];


if(isset($_POST['submit'])){

   $kname = $_POST['kname'];
   $kname = filter_var($kname, FILTER_SANITIZE_STRING);
   $pname = $_POST['pname'];
   $pname = filter_var($pname, FILTER_SANITIZE_STRING);
   $cat = $_POST['cat'];
   $cat = filter_var($cat, FILTER_SANITIZE_STRING);
   $price = $_POST['price'];
   $price = filter_var($price, FILTER_SANITIZE_STRING);
  
   $image = $_FILES['image']['name'];
   $image = filter_var($image, FILTER_SANITIZE_STRING);
   $image_size = $_FILES['image']['size'];
   $image_tmp_name = $_FILES['image']['tmp_name'];
   $image_folder = 'uploaded_img/'.$image;

   $insert = $conn->prepare("INSERT INTO `products`(pname,kname,category,price, image) VALUES(?,?,?,?,?)");
   $insert->execute([$pname,$kname,$cat, $price, $image]);

   if($insert){
        if($image_size > 2000000){
            $message[] = 'image size is too large!';
        }else{
            move_uploaded_file($image_tmp_name, $image_folder);
            $message[] = 'Uploaded successfully!';
            header('location:home.php');
        }
    }
   $select = $conn->prepare("SELECT * FROM `kitchens` WHERE kname = ?");
   $select->execute([$kname]);

   if($select->rowCount() > 0){
      $message[] = 'Kitchen name already exist!';
   }else{
      $insert_kit = $conn->prepare("INSERT INTO `kitchens`(user_id,kname) VALUES(?,?)");
      $insert_kit->execute([$user_id,$kname]);
   }
}

?>

<!DOCTYPE html>
<html lang="en">
<head>
   <meta charset="UTF-8">
   <meta http-equiv="X-UA-Compatible" content="IE=edge">
   <meta name="viewport" content="width=device-width, initial-scale=1.0">
   <title>Upload</title>

   <!-- font awesome cdn link  -->
   <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.1.1/css/all.min.css">

   <!-- custom css file link  -->
   <link rel="stylesheet" href="css/home.css">
   <link rel="stylesheet" href="css/sidenav.css">
   <link rel="stylesheet" href="css/style4.css">

</head>
<body>

<?php

if(isset($message)){
   foreach($message as $message){
      echo '
      <div class="message">
         <span>'.$message.'</span>
         <i class="fas fa-times" onclick="this.parentElement.remove();"></i>
      </div>
      ';
   }
}
include 'header.php';
?>
  
<section class="form-container">

   <form action="" enctype="multipart/form-data" method="POST">
      <h3>Upload new food</h3>
      <input type="text" name="kname" class="box" placeholder="Enter your kitchen name" required>
      <input type="text" name="pname" class="box" placeholder="Product name" required>
      <select class="box" name="cat">
         <?php 
         $select_cat = $conn->prepare("SELECT * FROM `category_list`");
         $select_cat->execute();
         if($select_cat->rowCount() > 0){
            while($fetch_cat = $select_cat->fetch(PDO::FETCH_ASSOC)){ 
      ?>
      <option value="<?php echo $fetch_cat["id"];?>">
                    <?php echo $fetch_cat["name"];?>
      </option>
      <?php 
            }
         }else{
            echo '<p class="empty">No categories added yet!</p>';
         }
      ?>
      </select>
      
      <input type="number" name="price" class="box" placeholder="Price" required>
      <input type="file" name="image" class="box" required accept="image/jpg, image/jpeg, image/png">
      <input type="submit" value="Upload now" class="btn" name="submit">
   </form>

</section>
<?php include 'footer.php'; ?>
<script>
    const menuToggle=document.querySelector('.menuToggle');
    const navigation=document.querySelector('.navigation');
    menuToggle.onclick=function(){
        navigation.classList.toggle('open');
    }
</script>

</body>
</html>