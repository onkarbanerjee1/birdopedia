<html>
    <head>
        <title>Add a bird</title>
    </head>
    <body>
        <form action="/addBirds" method="POST">
        Generic Name : <input type="text" name="name">
        <br><br><br>
        Common Name : <input type="text" name="common_name">
        <br><br><br>
        Scientific Name : <input type="text" name="scientific_name">
        <br><br><br>
        <input type="submit" value="Add Bird">
        </form>
    </body>
</html>