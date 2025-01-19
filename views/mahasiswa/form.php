<div class="modal-dialog" role="document">
    <div class="modal-content">
        <div class="modal-header">
            <h5 class="modal-title" id="exampleModalLabel">{{ .title }}</h5>
            <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                <span aria-hidden="true">&times;</span>
            </button>
        </div>
        <form id="form-mahasiswa" action="/mahasiswa/store" method="POST">
            <div class="modal-body">
                <div class="form-group">
                    <label for="">Nama Lengkap</label>
                    <input type="text" name="nama" value="{{ .mahasiswa.Nama }}" class="form-control">
                </div>
                <div class="form-group">
                    <label for="">Program Studi</label>
                    <select name="prodi" class="form-control">
                        <option {{ if eq .mahasiswa.Prodi `Sistem Informasi` }} selected {{ end }} value="Sistem Informasi">Sistem Informasi</option>
                        <option {{ if eq .mahasiswa.Prodi `Teknik Informatika` }} selected {{ end }} value="Teknik Informatika">Teknik Informatika</option>
                        <option {{ if eq .mahasiswa.Prodi `Teknologi Informasi` }} selected {{ end }} value="Teknologi Informasi">Teknologi Informasi</option>
                    </select>
                </div>
                <div class="form-group">
                    <label for="">Kelas</label>
                    <select name="kelas" class="form-control">
                        <option {{ if eq .mahasiswa.Kelas `A` }} selected {{ end }} value="A">A</option>
                        <option {{ if eq .mahasiswa.Kelas `B` }} selected {{ end }} value="B">B</option>
                        <option {{ if eq .mahasiswa.Kelas `C` }} selected {{ end }} value="C">C</option>
                    </select>
                </div>
            </div>
            <div class="modal-footer">
                {{ if .mahasiswa.NPM }}
                    <input type="hidden" name="npm" value="{{ .mahasiswa.NPM }}">
                    <button type="button" class="btn btn-secondary" data-dismiss="modal">Close</button>
                    <button type="submit" class="btn btn-primary">UBAH DATA</button>
                {{ else }}
                    <button type="button" class="btn btn-secondary" data-dismiss="modal">Close</button>
                    <button type="submit" class="btn btn-primary">SIMPAN DATA</button>
                {{ end }}
            </div>
        </form>
    </div>
</div>