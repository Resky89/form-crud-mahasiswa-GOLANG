<!doctype html>
<html lang="en">
  <head>
    
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">

    
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@4.0.0/dist/css/bootstrap.min.css" integrity="sha384-Gn5384xqQ1aoWXA+058RXPxPg6fy4IWvTNh0E263XmFcJlSAwiGgFAW/dAiS6JXm" crossorigin="anonymous">

    <title>Mahasiswa</title>
  </head>
  <body>

    <div class="modal fade" id="modalMahasiswa" tabindex="-1" role="dialog" aria-labelledby="modalMahasiswaLabel" aria-hidden="true">
        
    </div>
    
    <div class="container mt-5">
         <div class="form-group">
        <input type="text" id="search-mahasiswa" class="form-control" placeholder="Cari berdasarkan nama...">
    </div>

        <button type="button" class="btn btn-primary add-mahasiswa">Tambah Data</button>

        <table class="table mt-3">
            <thead>
                <th>#</th>
                <th>Nama Lengkap</th>
                <th>Prodi</th>
                <th>Kelas</th>
                <th>Aksi</th>
            </thead>
            <tbody>
                {{ .data }}
            </tbody>
        </table>
    </div>

    <script src="https://cdnjs.cloudflare.com/ajax/libs/jquery/3.6.0/jquery.min.js"></script>
    <script src="https://cdn.jsdelivr.net/npm/popper.js@1.12.9/dist/umd/popper.min.js" integrity="sha384-ApNbgh9B+Y1QKtv3Rn7W3mgPxhU9K/ScQsAP7hUibX39j7fakFPskvXusvfa0b4Q" crossorigin="anonymous"></script>
    <script src="https://cdn.jsdelivr.net/npm/bootstrap@4.0.0/dist/js/bootstrap.min.js" integrity="sha384-JZR6Spejh4U02d8jOt6vLEHfe/JQGiRRSQQxSfFWpi1MquVdAyjUar5+76PVCmYl" crossorigin="anonymous"></script>
  
    <script>

        var site_url = "http://localhost:8000/";

        $(document).ready(function(){
            $('.add-mahasiswa').click(function(){
                $.get(site_url + "mahasiswa/get_form", function(html){
                    $('#modalMahasiswa').html(html).modal('show')
                });     
            });

            $(document).on('click', '.edit-mahasiswa', function(){
                var npm = $(this).attr('data-npm');
                $.get(site_url + "mahasiswa/get_form?npm=" + npm, function(html){
                    $('#modalMahasiswa').html(html).modal('show')
                });     
            });

            $(document).on('click', '.delete-mahasiswa', function(){
                var npm = $(this).attr('data-npm');
                var confirmDelete = confirm("Apakah anda yakin ingin menghapus data?");
                if (confirmDelete == true) {
                    $.post(site_url + "mahasiswa/delete", {npm: npm}, function(response){
                        alert(response.message)
                        $('tbody').html(response.data);
                    }, 'JSON');
                }
            });

            $(document).on('submit', '#form-mahasiswa', function(e){
                e.preventDefault();
                $.ajax({
                    type: $(this).attr('method'),
                    url: $(this).attr('action'),
                    data: $(this).serialize(),
                    dataType: "json",
                    success: function(response){
                        $('tbody').html(response.data)
                        alert(response.message)
                        $('#modalMahasiswa').modal('hide')
                    },
                    error: function(response){
                        console.log(response)
                    }
                })
            });
        });

        $('#search-mahasiswa').on('input', function(){
            var searchTerm = $(this).val().trim();
            if (searchTerm === '') {
                $('tbody').html(originalData);
            } else {
                $.get(site_url + "mahasiswa/search?search=" + searchTerm, function(html){
                    $('tbody').html(html);
                });
            }
        });

        var originalData = '';
        $(document).ready(function(){
           originalData = $('tbody').html();
        });

    </script>
    </body>
</html>