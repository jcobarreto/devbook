$('#unfollow').on('click', unfollow);
$('#follow').on('click', follow);

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
