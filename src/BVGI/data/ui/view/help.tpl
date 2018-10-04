{{define "content"}}
<div class="jumbotron mb-0 pb-5">
    <div class="container">
        <div class="col-xs-12 col-center-block mb-5">
            <h1 align="center" class="text-primary">Help</h1>
        </div>
        <div class="row">
            <div class="column">
                <table class="table col-sm-12 table-bordered table-hover .table-striped px-0 mx-0">
                    <thead class="thead-inverse col-sm-12 px-0 mx-0">
                        <tr class="row col-sm-12 px-0 mx-0">
                            <th class="col-sm-4 text-center">Chức năng</th>
                            <th class="col-sm-8 text-center">Hướng dẫn</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr class="row col-sm-12 table-active px-0 mx-0">
                            <td class="col-sm-4">
                                1. Thẻ @gender, @Gender
                            </td>
                            <td class="col-sm-8">
                                Nếu trong nội dung của câu trả lời có thẻ này, bot sẽ dựa vào giới tính của khách hàng để thay thế bằng từ xưng hô tương ứng. @gender sẽ được thay bằng từ viết thường: "anh", "chị", "quý khách". @Gender sẽ được thay bằng từ viết hoa: "Anh", "Chị", "Quý khách".
                            </td>
                        </tr>

                        <tr class="row col-sm-12 table-active px-0 mx-0">
                            <td class="col-sm-4">
                                2. Thẻ @full_name, @first_name, @last_name
                            </td>
                            <td class="col-sm-8">
                                Nếu trong nội dung của câu trả lời có thẻ này, bot sẽ dựa vào tên trên Facebook của khách hàng để xưng hô tên. Ví dụ: "Kính chào @gender @first_name" -> "Kính chào anh Minh".
                            </td>
                        </tr>

                        <tr class="row col-sm-12 table-active px-0 mx-0">
                            <td class="col-sm-4">
                                3. Thẻ @call_staff
                            </td>
                            <td class="col-sm-8">
                                Nếu trong nội dung của câu trả lời có thẻ này, ngoài việc trả lời bot sẽ nhắn tin báo nhân viên chăm sóc khách hàng vào hỗ trợ tiếp.
                            </td>
                        </tr>

                        <tr class="row col-sm-12 table-active px-0 mx-0">
                            <td class="col-sm-4">
                                4. /restart
                            </td>
                            <td class="col-sm-8">
                                Khi chat với bot, nếu muốn bắt đầu lại cuộc hội thoại thì gõ lệnh này. Thường dùng lệnh này để test.
                            </td>
                        </tr>

                        <tr class="row col-sm-12 table-active px-0 mx-0">
                            <td class="col-sm-4">
                                5. @image[địa_chỉ_ảnh]
                            </td>
                            <td class="col-sm-8">
                                Nếu trong nội dung của câu trả lời có thẻ này, bot sẽ chuyển "@image[địa_chỉ_ảnh]" thành ảnh.
                            </td>
                        </tr>

                    </tbody>
                </table>
            </div>
        </div>
    </div>
</div>
{{end}}