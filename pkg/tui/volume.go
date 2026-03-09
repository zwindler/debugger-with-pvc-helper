package tui

import (
	"fmt"
	"strings"
)

// volumeView renders the volume multi-select step.
func volumeView(m model) string {
	var b strings.Builder

	b.WriteString(titleStyle.Render("Select PVC Volumes to Mount"))
	b.WriteString("\n")
	b.WriteString(subtitleStyle.Render(fmt.Sprintf("Pod: %s/%s", m.selectedNamespace, m.selectedPod.Name)))
	b.WriteString("\n\n")

	if len(m.selectedPod.PVCVolumes) == 0 {
		b.WriteString(dimStyle.Render("  No PVC volumes found on this pod"))
		b.WriteString("\n")
	} else {
		for i, vol := range m.selectedPod.PVCVolumes {
			cursor := "  "
			if i == m.cursor {
				cursor = cursorStyle.Render("> ")
			}

			checkbox := checkboxUnchecked
			if m.volumeSelected[i] {
				checkbox = checkboxChecked
			}

			name := vol.VolumeName
			if i == m.cursor {
				name = selectedStyle.Render(vol.VolumeName)
			}

			// Build detail string
			details := []string{vol.ClaimName}
			if vol.Size != "" {
				details = append(details, vol.Size)
			}
			if vol.AccessModes != "" {
				details = append(details, vol.AccessModes)
			}
			if vol.StorageClass != "" {
				details = append(details, vol.StorageClass)
			}
			if vol.MountPath != "" {
				details = append(details, fmt.Sprintf("mounted at %s", vol.MountPath))
			}

			detail := dimStyle.Render(fmt.Sprintf(" (%s)", strings.Join(details, ", ")))
			fmt.Fprintf(&b, "%s%s %s%s\n", cursor, checkbox, name, detail)
		}
	}

	// Show selection count
	selected := 0
	for _, v := range m.volumeSelected {
		if v {
			selected++
		}
	}

	b.WriteString("\n")
	if selected > 0 {
		b.WriteString(successStyle.Render(fmt.Sprintf("  %d volume(s) selected", selected)))
	} else {
		b.WriteString(warningStyle.Render("  No volumes selected yet"))
	}

	b.WriteString(helpStyle.Render("\n\n  j/k or arrows: navigate | space: toggle | enter: confirm | esc: back | q: quit"))

	return b.String()
}
