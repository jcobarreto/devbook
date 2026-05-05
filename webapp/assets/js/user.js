$('#unfollow').on('click', unfollow);
$('#follow').on('click', follow);
$('#edit-user').on('submit', edit);
$('#update-password').on('submit', updatePassword);
$('#delete-user').on('click', deleteUser);

function unfollow() {
  const UsedId = $(this).data('user-id');
  $(this).prop('disabled', true);

  $.ajax({
    url: `/users/${UsedId}/unfollow`,
    method: "POST",
  }).done(function() {
    window.location = `/users/${UsedId}`;
  }).fail(function() {
    Swal.fire("Ops...", "Error unfollowing user!", "error");
    $('#unfollow').prop('disabled', false);
  });
}

function follow() {
  const UsedId = $(this).data('user-id');
  $(this).prop('disabled', true);

  $.ajax({
    url: `/users/${UsedId}/follow`,
    method: "POST",
  }).done(function() {
    window.location = `/users/${UsedId}`;
  }).fail(function() {
    Swal.fire("Ops...", "Error following user!", "error");
    $('#follow').prop('disabled', false);
  });
}

function edit(event) {
  event.preventDefault();

  $.ajax({
    url: '/edit-user',
    method: "PUT",
    data: {
      name: $('#name').val(),
      email: $('#email').val(),
      nick: $('#nick').val(),
    }
  }).done(function() {
    Swal.fire("Success!", "User updated successfully!", "success")
        .then(function() {
          window.location = '/profile';
      });
  }).fail(function() {
    Swal.fire("Ops...", "Error updating user!", "error");
  });
}

function updatePassword(event) {
  event.preventDefault();

  if ($('#new-password').val() !== $('#confirm-password').val()) {
    Swal.fire("Ops...", "New password and confirmation do not match!", "error");
    return;
  }
  console.log($('#current-password').val(), $('#new-password').val());
  $.ajax({
    url: '/update-password',
    method: "POST",
    data: {
      current: $('#current-password').val(),
      new: $('#new-password').val(),
    }
  }).done(function() {
    Swal.fire("Success!", "Password updated successfully!", "success")
        .then(function() {
          window.location = '/profile';
      });
  }).fail(function() {
    Swal.fire("Ops...", "Error updating password!", "error");
  });
}

function deleteUser() {
  Swal.fire({
    title: "Are you sure?",
    text: "This action cannot be undone!",
    icon: "warning",
    showCancelButton: true,
    confirmButtonText: "Yes, delete my account!",
    cancelButtonText: "No, keep my account!"
  }).then((result) => {
    if (result.isConfirmed) {
      $.ajax({
        url: "/delete-user",
        method: "DELETE",
      }).done(function() {
        Swal.fire("Deleted!", "Your account has been deleted.", "success")
            .then(function() {
              window.location = '/logout';
            });
      }).fail(function() {
        Swal.fire("Ops...", "Error deleting user!", "error");
      });
    }
  });
}
