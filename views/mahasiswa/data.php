{{ $nomor := 0 }}
{{ range .mahasiswa }}
    <tr>
        {{ $nomor = increment $nomor 1 }}
        <td>{{ $nomor }}</td>
        <td>{{ .Nama }}</td>
        <td>{{ .Prodi }}</td>
        <td>{{ .Kelas }}</td>
        <td>
            <button data-npm="{{ .NPM }}" class="btn btn-danger btn-sm delete-mahasiswa">Delete</button>
            <button data-npm="{{ .NPM }}" class="btn btn-warning btn-sm edit-mahasiswa">Edit</button>
        </td>
    </tr>
{{ end }}