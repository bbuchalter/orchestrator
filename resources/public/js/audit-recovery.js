$(document).ready(function() {
  $("#audit_recovery_steps").hide();
  showLoader();
  var apiUri = "/api/audit-recovery/" + currentPage();
  if (auditCluster()) {
    apiUri = "/api/audit-recovery/cluster/" + auditCluster() + "/" + currentPage();
  }
  if (recoveryId() > 0) {
    apiUri = "/api/audit-recovery/id/" + recoveryId();
  }
  if (recoveryUid() != "") {
    apiUri = "/api/audit-recovery/uid/" + recoveryUid();
  }
  $.get(appUrl(apiUri), function(auditEntries) {
    auditEntries = auditEntries || [];
    displayAudit(auditEntries);
  }, "json");

  function auditInfo(audit) {
    var moreInfo = "";
    if (audit.Acknowledged) {
      moreInfo += '<div>Acknowledged by ' + audit.AcknowledgedBy + ', ' + audit.AcknowledgedAt + '<ul>';
      moreInfo += "<li>" + audit.AcknowledgedComment + "</li>";
      moreInfo += '</ul></div>';
    } else {
      moreInfo += '<div><strong>Unacknowledged</strong></div>';
    }
    if (audit.LostReplicas.length > 0) {
      moreInfo += "<div>Lost replicas:<ul>";
      audit.LostReplicas.forEach(function(instanceKey) {
        moreInfo += "<li><code>" + getInstanceTitle(instanceKey.Hostname, instanceKey.Port) + "</code></li>";
      });
      moreInfo += "</ul></div>";
    }
    if (audit.ParticipatingInstanceKeys.length > 0) {
      moreInfo += "<div>Participating instances:<ul>";
      audit.ParticipatingInstanceKeys.forEach(function(instanceKey) {
        moreInfo += "<li><code>" + getInstanceTitle(instanceKey.Hostname, instanceKey.Port) + "</code></li>";
      });
      moreInfo += "</ul></div>";
    }
    if (audit.AnalysisEntry.SlaveHosts.length > 0) {
      moreInfo += '<div>' + audit.AnalysisEntry.CountReplicas + ' replicating hosts :<ul>';
      audit.AnalysisEntry.SlaveHosts.forEach(function(instanceKey) {
        moreInfo += "<li><code>" + getInstanceTitle(instanceKey.Hostname, instanceKey.Port) + "</code></li>";
      });
      moreInfo += "</ul></div>";
    }
    if (audit.AllErrors.length > 0 && audit.AllErrors[0]) {
      moreInfo += "All errors:<ul>";
      audit.AllErrors.forEach(function(err) {
        moreInfo += "<li>" + err;
      });
      moreInfo += "</ul>";
    }
    moreInfo += '<div><a href="' + appUrl('/web/audit-failure-detection/id/' + audit.LastDetectionId) + '">Related detection</a></div>';
    moreInfo += '<div>Proccessed by <code>' + audit.ProcessingNodeHostname + '</code></div>';
    return moreInfo;
  }
  function displaySingleAudit(audit) {
    $("#audit .pager").hide();
    $("#audit_recovery_table").hide();

    $("#audit_recovery_details thead h3").text(audit.AnalysisEntry.Analysis)

    var appendRow = function(td1, td2) {
      var row = $('<tr/>');
      $('<td/>', {
        text: td1
      }).appendTo(row);
      $('<td/>', {
        html: td2
      }).appendTo(row);

      row.appendTo($("#audit_recovery_details tbody"));
    }
    appendRow("Failed instance", getInstanceTitle(audit.AnalysisEntry.AnalyzedInstanceKey.Hostname, audit.AnalysisEntry.AnalyzedInstanceKey.Port))
    var successor = getInstanceTitle(audit.SuccessorKey.Hostname, audit.SuccessorKey.Port);
    if (!audit.IsSuccessful) {
      successor = '<span class="text-danger"><span class="glyphicon glyphicon-remove-sign"></span> FAIL '+successor+'</span>';
    } else {
      successor = '<span class="text-success"><span class="glyphicon glyphicon-ok-sign"></span> '+successor+'</span>';
    }
    appendRow("Successor", successor)
    var clusterAlias = audit.AnalysisEntry.ClusterDetails.ClusterAlias;
    var clusterName = audit.AnalysisEntry.ClusterDetails.ClusterName;
    if (clusterAlias != clusterName) {
      appendRow("Cluster alias", '<a href="/web/cluster/alias/'+clusterAlias+'">' + clusterAlias + '</a>')
    }
    appendRow("Cluster name", '<a href="/web/cluster/'+clusterName+'">' + clusterName + '</a>')
    appendRow("Affected replicas", audit.AnalysisEntry.CountReplicas)
    appendRow("Start time", audit.RecoveryStartTimestamp)
    appendRow("End time", audit.RecoveryEndTimestamp)

    var numRows = $("#audit_recovery_details tbody tr").length;
    $('<td/>', {
      html: auditInfo(audit)
    }).attr("rowspan", numRows).appendTo($("#audit_recovery_details tbody tr:first-child"));

    auditRecoverySteps(audit.UID, $('#audit_recovery_steps'))
    $("#audit_recovery_steps").show();
  }

  function displayAudit(auditEntries) {
    var baseWebUri = appUrl("/web/audit-recovery/");
    if (auditCluster()) {
      baseWebUri += "cluster/" + auditCluster() + "/";
    }
    var singleRecoveryAudit = (auditEntries.length == 1);

    hideLoader();
    auditEntries.forEach(function(audit) {
      if (singleRecoveryAudit) {
        displaySingleAudit(audit)
        return;
      }

      var analyzedInstanceDisplay = getInstanceTitle(audit.AnalysisEntry.AnalyzedInstanceKey.Hostname, audit.AnalysisEntry.AnalyzedInstanceKey.Port);
      var sucessorInstanceDisplay = getInstanceTitle(audit.SuccessorKey.Hostname, audit.SuccessorKey.Port);
      var row = $('<tr/>');
      var ack = $('<span class="pull-left glyphicon acknowledge-indicator" title=""></span>');
      if (audit.Acknowledged) {
        ack.addClass("glyphicon-ok-sign").addClass("text-primary");
        var ackTitle = "Acknowledged by " + audit.AcknowledgedBy + " at " + audit.AcknowledgedAt + ": " + audit.AcknowledgedComment;
        ack.attr("title", ackTitle);
      } else {
        ack.addClass("glyphicon-question-sign").addClass("text-danger").addClass("unacknowledged");
        ack.attr("data-recovery-id", audit.Id);
        ack.attr("title", "Unacknowledged. Click to acknowledge");
      }
      var moreInfoElement = $('<a href="' + appUrl('/web/audit-recovery/uid/' + audit.UID) + '"><span class="pull-right glyphicon glyphicon-info-sign text-primary" title="More info"></span></a>');
      moreInfoElement.attr("data-recovery-id", audit.Id);

      var analysisTd = $('<td/>', {
        text: audit.AnalysisEntry.Analysis
      }).prepend(ack);
      if (!singleRecoveryAudit) {
        analysisTd.prepend(moreInfoElement)
      }
      analysisTd.appendTo(row);
      $('<a/>', {
        text: analyzedInstanceDisplay,
        href: appUrl("/web/search/" + analyzedInstanceDisplay)
      }).wrap($("<td/>")).parent().appendTo(row);
      $('<td/>', {
        text: audit.AnalysisEntry.CountReplicas
      }).appendTo(row);
      $('<a/>', {
        text: audit.AnalysisEntry.ClusterDetails.ClusterName,
        href: appUrl("/web/cluster/" + audit.AnalysisEntry.ClusterDetails.ClusterName)
      }).wrap($("<td/>")).parent().appendTo(row);
      $('<a/>', {
        text: audit.AnalysisEntry.ClusterDetails.ClusterAlias,
        href: appUrl("/web/cluster/alias/" + audit.AnalysisEntry.ClusterDetails.ClusterAlias)
      }).wrap($("<td/>")).parent().appendTo(row);
      $('<td/>', {
        text: audit.RecoveryStartTimestamp
      }).appendTo(row);
      $('<td/>', {
        text: audit.RecoveryEndTimestamp
      }).appendTo(row);
      if (audit.RecoveryEndTimestamp && !audit.IsSuccessful && !audit.SuccessorKey.Hostname) {
        $('<td/>', {
          text: "FAIL"
        }).appendTo(row);
      } else if (audit.SuccessorKey.Hostname) {
        $('<a/>', {
          text: sucessorInstanceDisplay,
          href: appUrl("/web/search/" + sucessorInstanceDisplay)
        }).wrap($("<td/>")).parent().appendTo(row);
      } else {
        $('<td/>', {
          text: "pending"
        }).appendTo(row);
      }
      var moreInfo = auditInfo(audit);
      row.appendTo('#audit_recovery_table tbody');

      var row = $('<tr/>');
      row.addClass("more-info");
      row.attr("data-recovery-id-more-info", audit.Id);
      $('<td colspan="8"/>').append(moreInfo).appendTo(row);
      if (audit.Acknowledged) {
        row.hide()
      }
      row.appendTo('#audit_recovery_table tbody');
    });
    if (singleRecoveryAudit) {
      $("[data-recovery-id-more-info]").show();
    }
    if (currentPage() <= 0) {
      $("#audit .pager .previous").addClass("disabled");
    }
    if (auditEntries.length == 0) {
      $("#audit .pager .next").addClass("disabled");
    }
    $("#audit .pager .previous").not(".disabled").find("a").click(function() {
      window.location.href = baseWebUri + (currentPage() - 1);
    });
    $("#audit .pager .next").not(".disabled").find("a").click(function() {
      window.location.href = baseWebUri + (currentPage() + 1);
    });
    $("#audit .pager .disabled a").click(function() {
      return false;
    });
    $("body").on("click", ".acknowledge-indicator.unacknowledged", function(event) {
      var recoveryId = $(event.target).attr("data-recovery-id");
      bootbox.prompt({
        title: "Acknowledge recovery",
        placeholder: "comment",
        callback: function(result) {
          if (result !== null) {
            showLoader();
            $.get(appUrl("/api/ack-recovery/" + recoveryId + "?comment=" + encodeURIComponent(result)), function(operationResult) {
              hideLoader();
              if (operationResult.Code == "ERROR") {
                addAlert(operationResult.Message)
              } else {
                location.reload();
              }
            }, "json");
          }
        }
      });
    });
  }
});
