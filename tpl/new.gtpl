<html>
    <head>
        <title>Add a bird</title>
    </head>
    <body>
        <form action="/birds/add" method="POST">
        Name : <input type="text" name="CommonName">
        <br><br><br>
        Scientific Name : <input type="text" name="ScientificName">
        <br><br><br>
        Picture URL : <input type="text" name="PictureURL">
        <br><br><br>
        <input type="checkbox" name="Endangered" value="true"> Endangered? <br>
        <br><br><br>
        Habitats:<br>
        <input type="checkbox" name="Habitat" value="North America"> North America <br>
        <input type="checkbox" name="Habitat" value="South America"> South America <br>
        <br><br><br>
        Posted By : <input type="text" name="PostedBy">
        <input type="submit" value="Add Bird">
        </form>
    </body>
</html>